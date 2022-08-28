contexts=${contexts:=/home/shy/contexts}

cd $contexts/usr/local/daemon/10000 && ./sbin/nginx -p $PWD
cd $contexts/usr/install/docker && dockerd --host unix://$PWD/docker.sock --pidfile $PWD/docker.pid --exec-root=$PWD/exec --data-root=$PWD/data --registry-mirror http://hub-mirror.c.163.com

su - shy -c "cd $contexts && ./bin/ice.bin forever start &"
su - shy -c "cd $contexts/usr/local/work/20220815-repos-server/ && ./bin/ice.bin forever start dev shy port 9030 nodename repos-server &"

