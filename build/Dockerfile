FROM ubuntu:rolling
#FROM debian:bullseye
# FROM balenalib/raspberry-pi-debian:latest
WORKDIR /opt/openvpn-admin-plus
EXPOSE 8080

#RUN apt-get update && apt-get install -y easy-rsa
#RUN chmod 755 /usr/share/easy-rsa/*
ADD assets/start.sh /opt/openvpn-admin-plus/start.sh
ADD assets/generate_ca_and_server_certs.sh /opt/openvpn-admin-plus/scripts/generate_ca_and_server_certs.sh
ADD assets/vars.template /opt/openvpn-admin-plus/scripts/

ADD openvpn-admin-plus.tar.gz /opt/openvpn-admin-plus/
RUN tar -zxf /opt/openvpn-admin-plus/openvpn-admin-plus.tar.gz --directory /opt/openvpn-admin-plus/
RUN rm -f /opt/openvpn-admin-plus/data.db
ADD assets/app.conf /opt/openvpn-admin-plus/conf/app.conf

CMD /opt/openvpn-admin-plus/start.sh
