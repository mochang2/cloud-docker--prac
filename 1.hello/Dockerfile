# Dockerfile is the origin file name for docker building
# If you want other names, you can use following instruction:
# docker build . -f thisismydockerfile

# https://nirsa.tistory.com/66?category=868315

# My Base Image
FROM ubuntu:20.04

# My Starting Point, instruction
ENTRYPOINT ["echo", "hello, world, bob"]

# docker build .
# docker build --tag myecho(:tag)
# docker rmi $(docker images -f "dangling=true" -q) (rm old-fashioned version)