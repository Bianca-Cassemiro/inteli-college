# Atividade Ponderada 2
A presente atividade demandou o desenvolvimento de um pacote no Robot Operating System (ROS), abordando as essenciais funcionalidades de mapeamento e navegação, utilizando o robô TurtleBot 3, seja em sua versão simulada ou real.
O repositório inclui um workspace ROS2 com, no mínimo, um pacote devidamente configurado, apresentando dois lançadores distintos: um destinado a inicializar todos os elementos necessários para o mapeamento e outro para a navegação. 
## Como utilizar 
1) Clone este repositório
2) Vá até a pasta src
3) Faça o build do pacotinho
   ```
   colcon build
   ```
4) Para realizar o mapeamento do local utilize:
   ```
    ros2 launch pacotinho test_launch.py
   ```
4) Para realizar a movimentação no local utilize:
   ```
    ros2 launch pacotinho move_launch.py
   ```
## Demonstração

[Screencast from 13-11-2023 13:08:48.webm](https://github.com/Bianca-Cassemiro/nav_pacotinho/assets/99203402/e14231a1-2d35-4abf-af3b-d93ab8f825f4)


Link do drive:
https://drive.google.com/file/d/1sj6lRObLu0dUkMtQmwKA3Qg2CxyG_zpm/view?usp=sharing

