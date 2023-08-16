cd /home/shy/docker && ./dockerd --host unix://$PWD/docker.sock --pidfile $PWD/docker.pid --exec-root=$PWD/exec --data-root=$PWD/data --registry-mirror "https://ccr.ccs.tencentyun.com" &
contexts=${contexts:=/home/shy/contexts}
cd $contexts/usr/local/daemon/10000 && ./sbin/nginx -p $PWD
su - shy -c "cd $contexts && ./bin/ice.bin forever start &"
