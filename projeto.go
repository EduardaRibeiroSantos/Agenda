package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	ARQ_LOGIN             = "login.txt"
	ARQ_AGENDA_TELEFONICA = "Agenda.txt"
	ARQ_TEMPORARIO        = "temp.txt"
)

type Aniversario struct {
	Dia, Mes, Ano int
}

type Telefone struct {
	DDD    int
	Numero uint
}

type Agenda struct {
	Registro int
	Nome     string
	Sexo     string
	Idade    int
	CPF      string
	Endereco string
	Tel      Telefone
	Niver    Aniversario
}

func main() {
	inicializacao()
}

func inicializacao() {
	validaLogin()
	menuPrincipal()
}

func validaLogin() {
	fmt.Println("Acesso restrito! Realize o LOGIN para continuar.")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\nUsuario: ")
	usuarioDigitado, _ := reader.ReadString('\n')
	usuarioDigitado = strings.TrimSpace(usuarioDigitado)

	fmt.Print("Senha: ")
	senhaDigitada, _ := reader.ReadString('\n')
	senhaDigitada = strings.TrimSpace(senhaDigitada)

	file, err := os.Open(ARQ_LOGIN)
	if err != nil {
		fmt.Println("Arquivo não pode ser aberto")
		return
	}
	defer file.Close()

	var usuarioArquivo, senhaArquivo string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linha := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(linha, "usuario:") {
			usuarioArquivo = strings.TrimSpace(strings.TrimPrefix(linha, "usuario:"))
		} else if strings.HasPrefix(linha, "senha:") {
			senhaArquivo = strings.TrimSpace(strings.TrimPrefix(linha, "senha:"))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}

	if usuarioDigitado == usuarioArquivo && senhaDigitada == senhaArquivo {
		fmt.Println("\nUsuario e senha CORRETOS")
	} else {
		fmt.Println("\nUsuario e senha INCORRETOS")
		fmt.Println("\nPressione Enter para prosseguir")
		reader.ReadString('\n')
		main()
	}
}

func menuPrincipal() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\t\t\tBEM VINDO AO SISTEMA DA AGENDA TELEFONICA")
		fmt.Println("DIGITE O NUMERO REFERENTE A OPCAO DESEJADA PARA MANIPULACAO DO SISTEMA")
		fmt.Println("1. Adicionar Cadastro \n2. Exibir Agenda \n3. Pesquisar Contato \n4. Modificar Contato \n5. Excluir contato \n6. Limpar a lista \n7. Sair")
		fmt.Print("Digite sua escolha: ")
		escolha, _ := reader.ReadString('\n')
		escolha = strings.TrimSpace(escolha)

		switch escolha {
		case "1":
			addContato()
		case "2":
			exibeAgenda()
		case "3":
			buscaContato()
		case "4":
			alteraContato()
		case "5":
			excluiContato()
		case "6":
			limparLista()
		case "7":
			fmt.Println("\nOBRIGADO POR USAR O SISTEMA")
			return
		default:
			fmt.Println("\nPor favor, insira apenas de 1 a 7.")
		}
	}
}

func addContato() {
	randGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Quantos cadastros deseja-se realizar? Digite 0 para cancelar")
	nStr, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nStr))
	if n == 0 {
		return
	}

	file, err := os.OpenFile(ARQ_AGENDA_TELEFONICA, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	for i := 0; i < n; i++ {
		var a Agenda
		a.Registro = randGenerator.Intn(100)
		fmt.Printf("\nNumero de registro gerado automaticamente: %d\n", a.Registro)
		fmt.Print("Insira o nome: ")
		a.Nome, _ = reader.ReadString('\n')
		a.Nome = strings.TrimSpace(a.Nome)

		fmt.Print("Insira o sexo (M)asculino (F)eminino: ")
		a.Sexo, _ = reader.ReadString('\n')
		a.Sexo = strings.TrimSpace(a.Sexo)

		fmt.Print("Data de nascimento DD MM AAAA: ")
		fmt.Scanf("%d %d %d\n", &a.Niver.Dia, &a.Niver.Mes, &a.Niver.Ano)
		a.Idade = time.Now().Year() - a.Niver.Ano
		fmt.Printf("Sua idade foi calculada automaticamente: %d anos\n", a.Idade)

		fmt.Print("Insira o CPF: ")
		a.CPF, _ = reader.ReadString('\n')
		a.CPF = strings.TrimSpace(a.CPF)

		fmt.Print("Insira o endereco: ")
		a.Endereco, _ = reader.ReadString('\n')
		a.Endereco = strings.TrimSpace(a.Endereco)

		fmt.Print("Insira o telefone DDD XXXXXXXXX: ")
		fmt.Scanf("%d %d\n", &a.Tel.DDD, &a.Tel.Numero)

		_, err := fmt.Fprintf(file, "%d|%s|%s|%d|%s|%s|%d|%d|%d|%d|%d\n",
			a.Registro, a.Nome, a.Sexo, a.Idade, a.CPF, a.Endereco,
			a.Tel.DDD, a.Tel.Numero, a.Niver.Dia, a.Niver.Mes, a.Niver.Ano)
		if err != nil {
			fmt.Println("Erro ao salvar o contato:", err)
		}

		fmt.Println("\nContato salvo")
		fmt.Println("\nPressione Enter para continuar...")
		reader.ReadString('\n')
	}
}

func exibeAgenda() {
	file, err := os.Open(ARQ_AGENDA_TELEFONICA)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var indicador int
	for scanner.Scan() {
		indicador++
		contato := strings.Split(scanner.Text(), "|")
		fmt.Printf("\nAs informacoes do contato [%d] sao:\n", indicador)
		fmt.Printf("\nRegistro: %s\nNome: %s\nSexo: %s\nData de nascimento: %s/%s/%s\nIdade: %s anos\nCPF: %s\nEndereco: %s\nTelefone: (%s) %s\n",
			contato[0], contato[1], contato[2], contato[8], contato[9], contato[10], contato[3], contato[4], contato[5], contato[6], contato[7])
		fmt.Println("\nInsira Enter para prosseguir para o proximo usuario da lista...")
		bufio.NewReader(os.Stdin).ReadString('\n')
	}
	fmt.Println("\nFim da lista! Adicione mais contatos para aparecerem aqui.")
	fmt.Println("\nInsira Enter para continuar...")
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func buscaContato() {
	reader := bufio.NewReader(os.Stdin)
	file, err := os.Open(ARQ_AGENDA_TELEFONICA)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	fmt.Print("\nVoce deseja pesquisar o contato pelo (n)ome ou (r)egistro? ")
	escolha, _ := reader.ReadString('\n')
	escolha = strings.TrimSpace(escolha)

	var nome string
	var registro int
	encontrado := false

	switch escolha {
	case "n", "N":
		fmt.Print("Digite o nome da pessoa a ser pesquisada: ")
		nome, _ = reader.ReadString('\n')
		nome = strings.TrimSpace(nome)
	case "r", "R":
		fmt.Print("Digite o registro gerado aleatoriamente da pessoa a ser pesquisada: ")
		fmt.Scanf("%d\n", &registro)
	default:
		fmt.Println("Opcao invalida.")
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contato := strings.Split(scanner.Text(), "|")
		if (escolha == "n" || escolha == "N") && contato[1] == nome {
			encontrado = true
		} else if (escolha == "r" || escolha == "R") && strconv.Itoa(registro) == contato[0] {
			encontrado = true
		}

		if encontrado {
			fmt.Printf("\nInformacoes detalhadas sobre %s\n", contato[1])
			fmt.Printf("\nRegistro: %s\nNome: %s\nSexo: %s\nData de nascimento: %s/%s/%s\nIdade: %s anos\nCPF: %s\nEndereco: %s\nTelefone: (%s) %s\n",
				contato[0], contato[1], contato[2], contato[8], contato[9], contato[10], contato[3], contato[4], contato[5], contato[6], contato[7])
			fmt.Println("\nInsira Enter para continuar...")
			reader.ReadString('\n')
			return
		}
	}

	if !encontrado {
		fmt.Println("\nContato nao encontrado.")
		fmt.Println("\nInsira Enter para continuar...")
		reader.ReadString('\n')
	}
}

func alteraContato() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite o registro do contato que deseja alterar: ")
	registroStr, _ := reader.ReadString('\n')
	registroStr = strings.TrimSpace(registroStr)
	registro, err := strconv.Atoi(registroStr)
	if err != nil {
		fmt.Println("Registro inválido.")
		return
	}

	file, err := os.Open(ARQ_AGENDA_TELEFONICA)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo de agenda:", err)
		return
	}
	defer file.Close()

	tempFile, err := os.Create(ARQ_TEMPORARIO)
	if err != nil {
		fmt.Println("Erro ao criar arquivo temporário:", err)
		return
	}
	defer tempFile.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contato := strings.Split(scanner.Text(), "|")
		contatoRegistro, _ := strconv.Atoi(contato[0])
		if contatoRegistro == registro {
			fmt.Println("Qual campo deseja modificar?")
			fmt.Println("1. Nome")
			fmt.Println("2. Sexo")
			fmt.Println("3. Data de Nascimento")
			fmt.Println("4. CPF")
			fmt.Println("5. Endereço")
			fmt.Println("6. Telefone")
			fmt.Print("Escolha a opção desejada: ")
			opcaoStr, _ := reader.ReadString('\n')
			opcaoStr = strings.TrimSpace(opcaoStr)
			opcao, err := strconv.Atoi(opcaoStr)
			if err != nil || opcao < 1 || opcao > 6 {
				fmt.Println("Opção inválida.")
				return
			}

			switch opcao {
			case 1:
				fmt.Print("Novo Nome: ")
				novoNome, _ := reader.ReadString('\n')
				contato[1] = strings.TrimSpace(novoNome)
			case 2:
				fmt.Print("Novo Sexo (M/F): ")
				novoSexo, _ := reader.ReadString('\n')
				contato[2] = strings.TrimSpace(novoSexo)
			case 3:
				fmt.Print("Nova Data de Nascimento (DD MM AAAA): ")
				novaData, _ := reader.ReadString('\n')
				novaData = strings.TrimSpace(novaData)
				dataSplit := strings.Split(novaData, " ")
				if len(dataSplit) != 3 {
					fmt.Println("Formato de data inválido.")
					return
				}
				dia, _ := strconv.Atoi(dataSplit[0])
				mes, _ := strconv.Atoi(dataSplit[1])
				ano, _ := strconv.Atoi(dataSplit[2])
				contato[8], contato[9], contato[10] = strconv.Itoa(dia), strconv.Itoa(mes), strconv.Itoa(ano)
			case 4:
				fmt.Print("Novo CPF: ")
				novoCPF, _ := reader.ReadString('\n')
				contato[4] = strings.TrimSpace(novoCPF)
			case 5:
				fmt.Print("Novo Endereço: ")
				novoEndereco, _ := reader.ReadString('\n')
				contato[5] = strings.TrimSpace(novoEndereco)
			case 6:
				fmt.Print("Novo Telefone (DDD XXXXXXXXX): ")
				novoTelefone, _ := reader.ReadString('\n')
				novoTelefone = strings.TrimSpace(novoTelefone)
				telefoneSplit := strings.Split(novoTelefone, " ")
				if len(telefoneSplit) != 2 {
					fmt.Println("Formato de telefone inválido.")
					return
				}
				ddd, _ := strconv.Atoi(telefoneSplit[0])
				numero, _ := strconv.Atoi(telefoneSplit[1])
				contato[6], contato[7] = strconv.Itoa(ddd), strconv.Itoa(numero)
			}
		}
		_, err := tempFile.WriteString(strings.Join(contato, "|") + "\n")
		if err != nil {
			fmt.Println("Erro ao escrever no arquivo temporário:", err)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao escanear o arquivo:", err)
		return
	}

	err = os.Rename(ARQ_TEMPORARIO, ARQ_AGENDA_TELEFONICA)
	if err != nil {
		fmt.Println("Erro ao renomear o arquivo temporário:", err)
		return
	}

	fmt.Println("Contato alterado com sucesso!")
}

func excluiContato() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite o registro do contato que deseja excluir: ")
	registroStr, _ := reader.ReadString('\n')
	registroStr = strings.TrimSpace(registroStr)
	registro, err := strconv.Atoi(registroStr)
	if err != nil {
		fmt.Println("Registro inválido.")
		return
	}

	file, err := os.Open(ARQ_AGENDA_TELEFONICA)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo de agenda:", err)
		return
	}
	defer file.Close()

	tempFile, err := os.Create(ARQ_TEMPORARIO)
	if err != nil {
		fmt.Println("Erro ao criar arquivo temporário:", err)
		return
	}
	defer tempFile.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contato := strings.Split(scanner.Text(), "|")
		contatoRegistro, _ := strconv.Atoi(contato[0])
		if contatoRegistro != registro {
			_, err := tempFile.WriteString(strings.Join(contato, "|") + "\n")
			if err != nil {
				fmt.Println("Erro ao escrever no arquivo temporário:", err)
				return
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao escanear o arquivo:", err)
		return
	}

	err = os.Rename(ARQ_TEMPORARIO, ARQ_AGENDA_TELEFONICA)
	if err != nil {
		fmt.Println("Erro ao renomear o arquivo temporário:", err)
		return
	}

	fmt.Println("Contato excluído com sucesso!")
}

func limparLista() {
	fmt.Println("Limpando a lista")

	arquivo, err := os.OpenFile(ARQ_AGENDA_TELEFONICA, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo de agenda:", err)
		os.Exit(1)
	}
	defer arquivo.Close()

	fmt.Println("Lista limpa com sucesso!")
}
