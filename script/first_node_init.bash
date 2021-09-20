#!/bin/bash
APP="blockchain"
APP_DAEMON=$APP"d"
APP_HOME=$HOME/.$APP

# go install
sudo rm -rf /usr/local/go
wget https://golang.org/dl/go1.17.linux-amd64.tar.gz    
sudo tar -xzf go1.17.linux-amd64.tar.gz -C /usr/local
rm -rf go1.17.linux-amd64.tar.gz
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
echo "export GOPATH=$GOPATH" > $HOME/.bash_profile
echo "export PATH=$PATH" >> $HOME/.bash_profile
mkdir $GOPATH
mkdir $GOPATH/bin
mkdir $GOPATH/src
mkdir $GOPATH/pkg

# blockchain app install
make
$APP_DAEMON keys add root --keyring-backend test
ROOT_ADDRESS=$($APP_DAEMON keys show root -a --keyring-backend test)
$APP_DAEMON init stock-chain --chain-id stock-chain
$APP_DAEMON add-genesis-account $ROOT_ADDRESS 110000000000stake
$APP_DAEMON gentx root 10000000000stake --chain-id stock-chain
$APP_DAEMON collect-gentxs
sed -i 's/enable = false/enable = true/' $APP_HOME/config/app.toml

# service register
sudo mkdir /var/log/$APP_DAEMON/
sudo touch /var/log/$APP_DAEMON/$APP_DAEMON.log
sudo touch /var/log/$APP_DAEMON/$APP_DAEMON-error.log
service="[Unit]
Description=$APP_DAEMON daemon
After=network-online.target
[Service]
User=$USER
ExecStart=$HOME/go/bin/$APP_DAEMON start --home=$APP_HOME
WorkingDirectory=$HOME/go/bin
Restart=always
RestartSec=3
LimitNOFILE=4096
[Install]
WantedBy=multi-user.target"
echo "$service" |& sudo tee /etc/systemd/system/$APP_DAEMON.service
sudo systemctl enable $APP_DAEMON
sudo systemctl start $APP_DAEMON

# check node
$HOME/go/bin/$APP_DAEMON --home=$APP_HOME tendermint show-node-id
journalctl -u $APP_DAEMON -f