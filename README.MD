# Разработка Интернет Приложений (РИП) 2023
## Бэкенд
### Вернуться на главную ([main](https://github.com/NikolayB800H/RestInPeace/tree/main))
### Запуск:
$
sudo service redis-server stop

cd ~/Projects/bmstu_5sem/RestInPeace/awesomeProject/

redis-server redis.conf --save 20 1 --loglevel warning --requirepass password

$
cd Projects/bmstu_5sem/RestInPeaceAsyncService/

. env/bin/activate

python weather_predictor/manage.py runserver 0.0.0.0:8000

$
sudo docker compose -f /home/nop/Projects/bmstu_5sem/RestInPeace/awesomeProject/docker-compose.yml up

---
Горкунов Николай ИУ5-53Б | Удаленные вычисления: 7. Прогноз погоды