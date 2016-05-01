from consul:latest

ADD https://releases.hashicorp.com/consul/0.6.4/consul_0.6.4_web_ui.zip /tmp/webui.zip
RUN mkdir -p /consul/ui && cd /tmp && unzip /tmp/webui.zip && rm /tmp/webui.zip && mv * /consul/ui

COPY default-config.json /consul/default/consul.json
