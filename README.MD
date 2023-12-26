(cd ~/swagger-editor/ && sudo docker run -d -p 81:8080 swagger-editor) # http://localhost:81/

sudo service redis-server stop
cd ~/Projects/bmstu_5sem/RestInPeace/awesomeProject/
redis-server redis.conf --save 20 1 --loglevel warning --requirepass password