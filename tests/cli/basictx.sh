#!/bin/bash

# these are two globals to control all scripts (can use eg. counter instead)
SERVER_EXE=basecoin
CLIENT_EXE=basecli

oneTimeSetUp() {
  # these are passed in as args
  BASE_DIR=$HOME/.basecoin_test_basictx
  CHAIN_ID=my-test-chain

  rm -rf $BASE_DIR 2>/dev/null
  mkdir -p $BASE_DIR


  # set up client - make sure you use the proper prefix if you set
  # a custom CLIENT_EXE
  export BC_HOME=${BASE_DIR}/client
  prepareClient

  # start basecoin server (with counter)
  initServer $BASE_DIR $CHAIN_ID 3456
  echo pid $PID_SERVER

  initClient $CHAIN_ID 3456

  echo "...Testing may begin!"
  echo
  echo
  echo
}

oneTimeTearDown() {
  echo
  echo
  echo "stopping $SERVER_EXE test server..."
  kill -9 $PID_SERVER >/dev/null 2>&1
  sleep 1
}

test00GetAccount() {
  SENDER=$(getAddr $RICH)
  RECV=$(getAddr $POOR)

  assertFalse "requires arg" "${CLIENT_EXE} query account"
  ACCT=$(${CLIENT_EXE} query account $SENDER)
  assertTrue "must have proper genesis account" $?
  assertEquals "no tx" "0" $(echo $ACCT | jq .data.sequence)
  assertEquals "has money" "9007199254740992" $(echo $ACCT | jq .data.coins[0].amount)

  ACCT2=$(${CLIENT_EXE} query account $RECV)
  assertFalse "has no genesis account" $?
}

test01SendTx() {
  SENDER=$(getAddr $RICH)
  RECV=$(getAddr $POOR)

  assertFalse "missing dest" "${CLIENT_EXE} tx send --amount=992mycoin --sequence=1 2>/dev/null"
  assertFalse "bad password" "echo foo | ${CLIENT_EXE} tx send --amount=992mycoin --sequence=1 --to=$RECV --name=$RICH 2>/dev/null"
  # we have to remove the password request from stdout, to just get the json
  RES=$(echo qwertyuiop | ${CLIENT_EXE} tx send --amount=992mycoin --sequence=1 --to=$RECV --name=$RICH 2>/dev/null | tail -n +2)
  assertTrue "sent tx" $?
  HASH=$(echo $RES | jq .hash | tr -d \")
  TX_HEIGHT=$(echo $RES | jq .height)
  assertEquals "good check" "0" $(echo $RES | jq .check_tx.code)
  assertEquals "good deliver" "0" $(echo $RES | jq .deliver_tx.code)

  # make sure sender goes down
  ACCT=$(${CLIENT_EXE} query account $SENDER)
  assertTrue "must have genesis account" $?
  assertEquals "one tx" "1" $(echo $ACCT | jq .data.sequence)
  assertEquals "has money" "9007199254740000" $(echo $ACCT | jq .data.coins[0].amount)

  # make sure recipient goes up
  ACCT2=$(${CLIENT_EXE} query account $RECV)
  assertTrue "must have new account" $?
  assertEquals "no tx" "0" $(echo $ACCT2 | jq .data.sequence)
  assertEquals "has money" "992" $(echo $ACCT2 | jq .data.coins[0].amount)

  # make sure tx is indexed
  TX=$(${CLIENT_EXE} query tx $HASH)
  assertTrue "found tx" $?
  assertEquals "proper height" $TX_HEIGHT $(echo $TX | jq .height)
  assertEquals "type=send" '"send"' $(echo $TX | jq .data.type)
  assertEquals "proper sender" "\"$SENDER\"" $(echo $TX | jq .data.data.inputs[0].address)
  assertEquals "proper out amount" "992" $(echo $TX | jq .data.data.outputs[0].coins[0].amount)
}


# load and run these tests with shunit2!
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )" #get this files directory

# load common helpers
. $DIR/common.sh

. $DIR/shunit2
