FROM node:12

RUN mkdir /workplace/
COPY gen/ /workplace/gen

RUN mkdir /workplace/tmp
RUN git clone https://github.com/michaelawyu/google-cloudevents /workplace/tmp
RUN mv /workplace/tmp/tools/quicktype-wrapper /workplace/gen/quicktype-wrapper

WORKDIR /workplace/gen/quicktype-wrapper
RUN npm install

WORKDIR /workplace/gen/
RUN npm install quicktype-wrapper/
RUN npm install
RUN npm link

CMD gen
