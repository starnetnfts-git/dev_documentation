# Starnetscan Explorer Build & Run

Note: the setup of the explorer is a very complicated process because of erlang / elixir and phoenix framework. If you are not familiar with them arm yourself with patience and about 6+ hours of troubleshooting

## Installing Blockscout dependencies

Follow the below command to install dependencies and continue from rust installation in the main
document.

```
sudo apt-get install git \
automake \
libtool inotify-tools \
libgmp-dev \
libgmp10 \
build-essential \
cmake -y
sudo apt-key add erlang_solutions.asc
sudo apt-get update
sudo apt-get install erlang -y
sudo apt-get install elixir=1.12.2-1 -y
sudo apt-get install nodejs -y
curl -sL https://deb.nodesource.com/setup_14.x | sudo -E bash -
sudo apt-get install -y nodejs
apt-get install cargo
```

## Setting Up the Application

If the dependencies failed to compile install the rebar dependency after copying the repo and
downloading the dependencies https://github.com/elixir-lang/elixir/issues/3857

```
● miix local.rebar
```

## Installing Postgres

```
● sudo sh -c 'echo "deb http://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg
main" > /etc/apt/sources.list.d/pgdg.list'
● wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key
add -
● sudo apt-get update
● apt-get install postgresql -y
```

## Initial Setup for postgres

```
● sudo su
● su - postgres
● psql
● ALTER USER postgres WITH PASSWORD ‘password’;
● \q
● exit
● exit
Note: replace password with the actual password in the command(in my case it’s 12345),
To check if you did everything correctly run “systemctl status postgresql”, you'll see active as active
```

## Installing Rust

```
● sudo su
● curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```

Enter 1 and let it install, after it’s finished start a new connection (shell) and close the current one, it
is needed so rust is added in your environment path. (this is important or application won't build)
Verify by command “rustup” if it is found then you did everything correctly, make sure to sudo su
before testing rustup as we installed rust for our root user.

## Setting Up the Application

● cd “path to directory where you want to keep the files”

```
● git clone https://github.com/poanetwork/blockscout
● cd blockscout
● mix deps.get
● mix phx.gen.secret
```

It may take time, it is needed for environment key for blackscout

### Save the secret string!!! it will be needed later

```
● mix do deps.get, local.rebar --force, deps.compile, compile
```

Replace MYPASSWORD with your postgress password in below command if it’s different

```
● export DATABASE_URL=postgresql://postgres:MYPASSWORD@localhost:5432/blockscout;
● mix do ecto.create, ecto.migrate
● cd apps/block_scout_web/assets; sudo npm install
● node_modules/webpack/bin/webpack.js --mode production
```

Wait ALOT! Go do something else and return in 15 minutes or something..

It may get stuck after 3 compilations, if not then you are good if it does, press Ctrl + C to exit out
. It’s a known webpack bug that won't affect the deployment.

```
● cd -
● cd apps/explorer
● sudo npm install
● cd -
● mix phx.digest
● cd apps/block_scout_web; mix phx.gen.cert blockscout blockscout.local; cd -
```

```
● nano /etc/hosts
```

and add blocksout and blockscout.local after 127.0.0.1 localhost

Update nginx config so the default port 80 point to 4000.

In the blockscout folder “sudo nano start.sh” and replace the postgres password and the secret key

```
#!/bin/bash
export DATABASE_URL=postgresql://postgres:xxxxxxxxxxxxxxxxxxxxxx@localhost:5432/blockscout;
export ETHEREUM_JSONRPC_HTTP_URL=http://localhost:8545
export ETHEREUM_JSONRPC_VARIANT=geth
export COIN=STAR
export LOGO=/images/k_logo.png
#export ETHEREUM_JSONRPC_WS_URL=ws://starnetscan.com/ws
export SECRET_KEY_BASE=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```

create the blockscout service

```
Description=Blockscout
After=network.target

[Service]
User=root
WorkingDirectory=/var/www/blockscout
ExecStart=/bin/bash /var/www/blockscout/start.sh
KillSignal=SIGHUP

[Install]
WantedBy=default.target
```

and that should be all.

If you reached this point, congratulate yourself!

## Extra: CSS Styling & Other

### Remove Market Value Graph:

File: blockscout/apps/block_scout_web/assets/css/components/\_dashboard-banner.scss
Add opacity:0; below padding in the above file

### Footer content:

File: blockscout/apps/block_scout_web/assets/css/components/\_footer.scss
.footer-body {
display: none;
}
Add the below in the above file at the end

### Adding the logo:

Location: blockscout/apps/block_scout_web/assets/static/images/
Drop your logo file in the above mentioned location and set the environment variable in the
startup.sh script in root blockscout folder

```
export LOGO=/images/custom_log.png
```

The app uses webpack to compile scss to actual css files, from the root blockscout folder run
the below command in terminal. It may take time

```
systemctl stop blockscout
cd apps/block_scout_web/assets; npm install &&
node_modules/webpack/bin/webpack.js --mode production; cd -
```

Once after the compilation generate the static assets. And restart the service

```
mix phx.digest
```

### Updating the symbol to star:

blockscout/apps/block_scout_web/priv/gettext/en/LC_MESSAGES/default.po
In the above find the below line and change the “msgstr” to “Star” to
update the symbol across the application

### Regenerating Static Assets (working Folder is blockscout)

Stop the service

```
systemctl stop blockscout
```

Delete existing static assets

```
mix phx.digest.clean
```

Make your css changes and run the below command to generate css from scss files

```
cd apps/block_scout_web/assets; npm install &&
node_modules/webpack/bin/webpack.js --mode production; cd -
```

Generate static assets using below command

```
mix phx.digest
```

Start the service

```
systemctl start blockscout
```

Repeat for the changes.

.dashboard-banner-network-graph

display: none
