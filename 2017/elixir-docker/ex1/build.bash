$ docker build -t parisex1 .
Sending build context to Docker daemon  3.072kB
Step 1/3 : FROM alpine:latest
 ---> 7328f6f8b418
Step 2/3 : COPY myscript.sh /myscript.sh
 ---> 8303d048f7de
Step 3/3 : ENTRYPOINT /myscript.sh
 ---> Running in 944bc2c734fa
 ---> f4a1632791a7
Removing intermediate container 944bc2c734fa
Successfully built f4a1632791a7
Successfully tagged parisex1:latest
