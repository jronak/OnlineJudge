#!bin/bash
container=judge1
if [ "$EUID" -ne 0 ]
  then echo "Please run as root"
  exit 1
fi
apt-get install lxc-dev
if [ "$?" -ne 0 ]
	then echo "*******Lxc package installation failed*********"
	exit 1
fi
echo "*********Installing primary container, please be patient***********"
lxc-create -t ubuntu -n $container
chmod a+rwx /var/lib/lxc/$container
chmod a+rwx /var/lib/lxc/$container/rootfs
chmod a+rxw /var/lib/lxc/$container/rootfs/home/ubuntu/
cp -R Judge /var/lib/lxc/$container/rootfs/home/ubuntu/
cp firstSetup.bash startJudge.bash restart.bash .lxcmain.go /var/lib/lxc/$container/rootfs/home/ubuntu/
echo "********Open another terminal, run the following command******
*lxc-start -n $container
*lxc-console -n $container
*Username:ubuntu
*Password:ubuntu
*Recommended to change the default Password
*Setup root Password
*As root run script firstSetup.bash*********"
exit 0
