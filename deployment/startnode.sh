if [ ! "$(docker ps -aq -f name=node)" ]; then
 # 请手动修改版本号
 docker run -d -p 37101:37101 -p 47101:47101 -p 37301:37301 --name node -v /work/matrixchain/data:/home/work/matrixchain/data -v /work/matrixchain/conf:/home/work/matrixchain/conf -v /work/matrixchain/logs:/home/work/matrixchain/logs superconsensuschain/matrixchain:v5.1.0
fi

if [ ! "$(docker ps -q -f name=node)" ]; then
	docker start node
	echo "start node finish"
fi
echo "start node finish"
