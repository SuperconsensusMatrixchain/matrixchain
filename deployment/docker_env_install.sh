#!/bin/bash
OS_VER=$( grep VERSION_ID /etc/os-release | cut -d'=' -f2 | sed 's/[^0-9\.]//gI' )
OS_MAJ=$(echo "${OS_VER}" | cut -d'.' -f1)
OS_MIN=$(echo "${OS_VER}" | cut -d'.' -f2)

MEM_MEG=$( free -m | sed -n 2p | tr -s ' ' | cut -d\  -f2 || cut -d' ' -f2 )
CPU_SPEED=$( lscpu | grep -m1 "MHz" | tr -s ' ' | cut -d\  -f3 || cut -d' ' -f3 | cut -d'.' -f1 )
CPU_CORE=$( nproc )
MEM_GIG=$(( ((MEM_MEG / 1000) / 2) ))
export JOBS=${JOBS:-$(( MEM_GIG > CPU_CORE ? CPU_CORE : MEM_GIG ))}

DISK_INSTALL=$(df -h . | tail -1 | tr -s ' ' | cut -d\  -f1 || cut -d' ' -f1)
DISK_TOTAL_KB=$(df . | tail -1 | awk '{print $2}')
DISK_AVAIL_KB=$(df . | tail -1 | awk '{print $4}')
DISK_TOTAL=$(( DISK_TOTAL_KB / 1048576 ))
DISK_AVAIL=$(( DISK_AVAIL_KB / 1048576 ))

printf "\\nOS name: ${OS_NAME}\\n"
printf "OS Version: ${OS_VER}\\n"
printf "CPU speed: ${CPU_SPEED}Mhz\\n"
printf "CPU cores: %s\\n" "${CPU_CORE}"
printf "Physical Memory: ${MEM_MEG} Mgb\\n"
printf "Disk install: ${DISK_INSTALL}\\n"
printf "Disk space total: ${DISK_TOTAL%.*}G\\n"
printf "Disk space available: ${DISK_AVAIL%.*}G\\n"

if [ "${MEM_MEG}" -lt 3000 ]; then
	printf "Your system must have 3 or more Gigabytes of physical memory installed.\\n"
	printf "Exiting now.\\n"
    printf "system memory configure failed.\\n"
    echo "install finish"
	exit 1
fi

if [ "${DISK_AVAIL%.*}" -lt 20 ]; then
	printf "You must have at least 20GB of available storage.\\n"
	printf "Exiting now.\\n"
    printf "system storage configure failed.\\n"
    echo "install finish"
	exit 1
fi
# 安装docker
command -v docker
if [ $? -ne 0 ]; then
    curl -fsSL https://get.docker.com -o get-docker.sh
    sudo sh get-docker.sh
fi
if [ `whoami` != "root" ]; then
    grep ^docker /etc/group >& /dev/null
    if [ $? -ne 0 ]; then
        sudo groupadd docker
        sudo usermod -aG docker ${USER}
    fi
fi
# 启动dokcer
sudo systemctl start docker

if [ $? -ne 0 ]; then
    yes | sudo apt-get purge docker-ce docker-ce-cli containerd.io
    sudo rm -rf /var/lib/docker
    sudo rm -rf /var/lib/containerd
    printf "docker install failed.\\n"
    echo "install finish"
    exit 1
else
    printf "docker install successful.\\n"
fi
# 配置matrixchain目录
mkdir -p /work/matrixchain/
cd /work/matrixchain

# 请手动修改版本号
# 拉取镜像
docker pull superconsensuschain/matrixchain:v5.1.0

if [ $? -ne 0 ]; then
    printf "docker pull failed.\\n"
    echo "install finish"
    exit 1
else
    printf "docker pull successful.\\n"
fi

if [ ! "$(docker ps -aq -f name=node)" ]; then
    # 请手动修改版本号
    docker run -it -d --rm --name node superconsensuschain/matrixchain:v5.1.0
    if [ $? -ne 0 ]; then
        printf "docker run failed.\\n"
        echo "install finish"
        exit 1
    else
        printf "docker run successful.\\n"
    fi
    sudo docker cp node:/home/work/matrixchain/data/ /work/matrixchain/
    sudo docker cp node:/home/work/matrixchain/conf/ /work/matrixchain/
    sudo docker cp node:/home/work/matrixchain/logs/ /work/matrixchain/
    docker stop node
    sudo rm /work/matrixchain/data/blockchain/* -rf
    sudo rm /work/matrixchain/data/netkeys -rf
    sudo rm /work/matrixchain/data/keys/* -rf
    sudo rm /work/matrixchain/logs/* -rf
    if [ $? -ne 0 ]; then
        printf "env pre failed.\\n"
        echo "install finish"
        exit 1
    else
        printf "env pre successful.\\n"
    fi
else
    printf "node already exist.\\n"
fi
