context=${context:=/home/shy/contexts}

cd $context/usr/local/daemon/10000 && ./sbin/nginx -p $PWD

su - shy -c "cd $context/ && ./bin/ice.bin forever start &"
su - shy -c "cd $context/usr/local/work/20220815-repos-server/ && ./bin/ice.bin forever start dev shy port 9030 nodename repos-server &"

