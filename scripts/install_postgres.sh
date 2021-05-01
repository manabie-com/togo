DEBIAN_FRONTEND=noninteractive apt-get install -y tzdata

apt-get install -y gpg
apt-get install -y wget
apt-get install -y apt-transport-https
apt-get install -y ca-certificates
apt-get install -y gpg-agent
#apt-get install -y lsb-release

wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -

#RELEASE=$(lsb_release -cs)
#echo $RELEASE

echo "deb http://apt.postgresql.org/pub/repos/apt/ bionic"-pgdg main | sudo tee  /etc/apt/sources.list.d/pgdg.list

# fetch the metadata from the new repo
sudo apt-get update

sudo apt-get install postgresql-11 -y
