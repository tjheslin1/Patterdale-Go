FROM golang:1.8.3

RUN mkdir -p $GOPATH/src/github.com/tjheslin1
&& cd $GOPATH/src/github.com/tjheslin1
&& git clone https://github.com/tjheslin1/Patterdale

# install required software
RUN apt-get update && apt-get install -y unzip vim pkg-config libaio1

RUN mkdir -p /usr/lib/oracle/12.1/client64/lib && mkdir -p /usr/include/oracle/12.1/client64/

ENV LD_LIBRARY_PATH /usr/lib:/usr/local/lib:/usr/lib/oracle/12.1/client64/lib
ENV NLS_LANG AMERICAN_AMERICA.AL32UTF8

COPY resources/instantclient-basic-linux.x64-12.1.0.2.0.zip /tmp/
COPY resources/instantclient-sdk-linux.x64-12.1.0.2.0.zip /tmp/

RUN cd /tmp && unzip instantclient-basic-linux.x64-12.1.0.2.0.zip
RUN cp /tmp/instantclient_12_1/* /usr/lib/oracle/12.1/client64/lib/

RUN cd /tmp && unzip instantclient-sdk-linux.x64-12.1.0.2.0.zip
RUN cp /tmp/instantclient_12_1/sdk/include/* /usr/include/oracle/12.1/client64/

RUN ln -s /usr/lib/oracle/12.1/client64/lib/libclntsh.so.12.1 /usr/lib/oracle/12.1/client64/lib/libclntsh.so
RUN ln -s /usr/lib/oracle/12.1/client64/lib/libocci.so.12.1 /usr/lib/oracle/12.1/client64/lib/libocci.so

RUN cp resources/oci8.pc /usr/lib/pkgconfig/oci8.pc

RUN rm -rf /tmp/*

# ENTRYPOINT ["go", "run", "main.go"]
CMD "/bin/bash"
