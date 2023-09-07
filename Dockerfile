FROM python:3.9

WORKDIR /usr/src

COPY . .

RUN pip install --upgrade pip

RUN pip install -r requirements.txt

EXPOSE 3000

 
