
if [ "$EUID" -ne 0 ]
  then echo "Please run as root"
  exit 1
fi
apt-get install gcc
if [ "$?" -ne 0 ]
	then echo "Gcc installation failed"
	exit 1
fi
apt-get install g++
if [ "$?" -ne 0 ]
	then echo "G++ installation failed"
	exit 1
fi
apt-get install python
if [ "$?" -ne 0 ]
	then echo "Python2 installation failed"
	exit 1
fi
apt-get install python3
if [ "$?" -ne 0 ]
	then echo "Python3 installation failed"
	exit 1
fi
apt-get install wget
wget https://storage.googleapis.com/golang/go1.4.2.linux-amd64.tar.gz
tar -C /usr/local -xzf go*.tar.gz
rm go*.tar.gz
apt-get install openjdk-6-jdk
if [ "$?" -ne 0 ]
	then echo "Python3 installation failed"
	exit 1
fi
i=1
while [ $i -lt 11 ]
do
	mkdir d$i
	chmod a+rwx d$i
	touch d$i/GoBackhomeKid
	chattr +i d$i/GoBackhomeKid
	((i=i+1))
done
mkdir -p go/src/Judge
chmod o-r go
export PATH=$PATH:/usr/local/go/bin
export GOPATH=/home/ubuntu/go
cp -R Judge firstSetup.bash *.bash go/src
rm *
rm -r Judge
cd go/src/Judge
cp main ../
rm main
go build
cd ..
go build main.go
cp main /home/ubuntu
cp startJudge.bash /home/ubuntu
cd /home/ubuntu
chmod o+rw main startJudge.bash
echo "Go to localhost:5000 Setup the constraints for the container
Then run the startJudge.bash as user
"