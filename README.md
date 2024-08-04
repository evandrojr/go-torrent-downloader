
# go-torrent-downloader

Este projeto é um servidor HTTP simples construído com Go, usando o framework Gin, que permite aos usuários iniciar downloads de torrents através de links magnet. Os downloads são gerenciados pelo cliente `aria2c`.

## Pré-requisitos

- **Go:** Certifique-se de ter o Go instalado em sua máquina. Você pode baixá-lo em [golang.org](https://golang.org/).
- **aria2c:** Este projeto utiliza o `aria2c` para gerenciar downloads de torrents. Instale-o usando o seguinte comando:
  
  ```bash
  sudo apt-get install aria2
  ```


Git: Para clonar o repositório, você precisará do Git. Instale-o usando o seguinte comando:

```bash
sudo apt-get install git
```
Configuração
Clone este repositório para sua máquina local:

```bash

git clone https://github.com/seu-usuario/torrent-downloader.git
cd torrent-downloader
```
Compile o programa:

```bash
go build
```

Certifique-se de que o diretório ~/Downloads existe e está acessível para o programa, pois é onde os arquivos serão salvos.

Como usar
Inicie o servidor:

Execute o binário gerado pelo comando go build:

bash
Copiar código
```
./torrent-downloader
```
O servidor estará disponível na porta 8091.

Envie um link magnet:

Você pode usar ferramentas como curl ou seu navegador para enviar um link magnet para o servidor.

Com o navegador:

Abra a página http://localhost:8091/ e use o formulário fornecido para enviar o link magnet.


Com curl:

```bash
curl -X POST -F "magnet=magnet:?xt=urn:btih:..." http://localhost:8091/download
```


Logs:

Os logs das operações do aria2c serão salvos em ./public/log.txt.

Estrutura do Código
main.go: O arquivo principal que configura o servidor e define os endpoints.

/public: Diretório onde os arquivos de log são armazenados e a página HTML é servida.

Segurança
Nota de Segurança: Este servidor executa comandos shell diretamente a partir de entrada do usuário. Em um ambiente de produção, é importante validar e sanitizar entradas do usuário para evitar possíveis ataques de injeção de comando.

Contribuições
Contribuições são bem-vindas! Se você encontrar algum problema ou tiver sugestões de melhorias, sinta-se à vontade para abrir uma issue ou enviar um pull request.

Licença
Este projeto está sob a Licença MIT.


### Considerações Adicionais

- **Segurança:** O código atual executa comandos de shell diretamente com base na entrada do usuário. Em ambientes de produção, é crucial implementar validações rigorosas para evitar vulnerabilidades de injeção de comandos.

- **Logs:** Verifique os logs para entender se algum problema ocorreu durante o download.

- **Diretórios:** Você pode modificar os caminhos de download e log de acordo com sua estrutura de diretórios preferida. 

Sinta-se à vontade para adaptar este `README.md` conforme suas necessidades específicas e o contexto do seu projeto!





