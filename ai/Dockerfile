FROM python:3.9

WORKDIR /app

COPY requirements.txt .

RUN pip install -r requirements.txt

RUN if [ "$PRODUCTION" = "1" ];then wget $MODEL_URL -O $ML_MODEL;fi

COPY . .

ENTRYPOINT [ "/app/entry.sh" ]
