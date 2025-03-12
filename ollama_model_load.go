package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"github.com/gookit/color"
)

var (
	red    = color.New(color.FgRed)
	green  = color.New(color.FgGreen)
	yellow = color.New(color.FgYellow)
	// miku_blue = 74cccf
	// miku_pink = ef669b
	miku_blue = color.HEX("#74cccf")
	miku_pink = color.HEX("#ef669b")
)

func init() {
	miku_blue.Println(`                                                                                            $         `)                  
	miku_pink.Print(`                                                                            `);yellow.Print(`$`);miku_blue.Println(`              $$$         `)                  
	miku_blue.Print(`                                                                           `);yellow.Print(`$$$$`);miku_blue.Println(`           $$$$$$      `)                   
	miku_blue.Print(`                                                                          `);yellow.Print(`$$$$$$`);miku_blue.Println(`     $$$$$$$$$$$$$$$$ `)                   
	miku_blue.Print(`                                                   $$$$$$$$$$$$$$$$$$    `);yellow.Print(`$$$$$$$$`);miku_blue.Println(`    $$$$$$"""$$$$$$$ `)                   
	miku_blue.Print(`                                             $$$$$$$$$$$$$$$$$$$$$$$$ `);yellow.Print(`$$$$$$..$$$$`);miku_blue.Println(`    $$$$$$$$$$$$$$$  `)                  
	miku_blue.Print(`                                          $$$$$$$$$$$$$$$$$$$$$$$$   `);yellow.Print(`$$$$$j....j$$$$$$$`);miku_blue.Print(`  $$$$$$`);miku_pink.Println(`    $$$$$$$$$$$$$$$$$     `)
	miku_blue.Print(`                                       $$$$$$$""$$$$$$$$$            `);yellow.Print(`$$$$$$j..%$$$$$$$$`);miku_blue.Print(`   $$$`);miku_pink.Println(`$$$$$$$$$$$$$$$$$$$$$$$$$   `)
	miku_blue.Print(`                                     $$$$$"""$$$$$`);miku_pink.Print(`$$`);miku_blue.Print(`         $$$$        `);yellow.Print(`$$$$$$$$$`);miku_pink.Println(`    $$$$$$$$$$$$jj((((((((((((((j$$$$$ `)
	miku_blue.Print(`                                   $$$$$"""$$$$$$`);miku_pink.Print(`$$$$`);miku_blue.Print(`        $$$$$$$      `);yellow.Print(`$$$$$$`);miku_pink.Println(`  $$$$$$$$$rj(((((((((((((((((((((((($$$$`)
	miku_blue.Print(`                                 $$$$$"""$$$$$  `);miku_pink.Print(`$$$$$$`);miku_blue.Print(`      $$$$$$$$$$     `);yellow.Print(`$$$`);miku_pink.Println(`$$$$$$$$jjjjj$$$$$$$$$$$$$$$$j(((((((((j$$$`)
	miku_blue.Print(`                               $$$$$""")$$$$ `);miku_pink.Print(`$$$$$$j$$$$$$`);miku_blue.Print(` $$$$""""$$$$$`);miku_pink.Println(`$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$j(((((($$$$`)
	miku_blue.Print(`                              $$$$""""$$$$$`);miku_pink.Print(`$$$$$$(((j$$$$$$`);miku_blue.Print(`$$$""""""""$$$$$`);miku_pink.Println(`$$$$$$$$$$$$$$                  $$$$j(((($$$$ `)
	miku_blue.Print(`                            $$$$$""""$$$$ `);miku_pink.Print(`$$$$$$$j($$$$$$$`);miku_blue.Print(`$$$$""""""""""p$$$$$`);miku_pink.Println(`$$$$$                          $$$((($$$$  `)
	miku_blue.Print(`                            $$$"""""$$$$      `);miku_pink.Print(`$$$$$$$$`);miku_blue.Print(`   $$$$""""""""""""""$$$$$`);miku_pink.Println(`                             $$$(($$$$   `)
	miku_blue.Print(`                           $$$"""""$$$$         `);miku_pink.Print(`$$$$$$  $`);miku_blue.Print(`$$$$""""""""""""""""$$$$$`);miku_pink.Println(`                          $$$$j$$$$    `)
	miku_blue.Print(`                          $$$$""""$$$$           `);miku_pink.Print(`$$$  $$$`);miku_blue.Print(`$$$""""""""""""""""""$$$$`);miku_pink.Println(`                       $$$$$$$$$$     `)
	miku_blue.Print(`                          $$$"""""$$$                `);miku_pink.Print(`$$$`);miku_blue.Print(`$$$$""""""""""""""""""""$$$$$`);miku_pink.Println(`                  $$$$$$$$$$$       `)
	miku_blue.Print(`                         $$$"""""$$$$               `);miku_pink.Print(`$$$`);miku_blue.Print(` $$$$""""""""""""""""""""""$$$$`);miku_pink.Println(`               $$$$$$$$$$          `)
	miku_blue.Print(`                         $$$"""""$$$                    $$$""""""""""""""""""""""""$$$$`);miku_pink.Println(`              $$$$$$$             `)
	miku_blue.Print(`                        $$$$"""""$$$                    $$$"""""""""""""""""""""""""$$$$`);miku_pink.Println(`             $$$$                `)
	miku_blue.Print(`                        $$$"""""$$$$                   $$$"""""""""""""""""""""""""""$$$$`);miku_pink.Println(`                                `)
	miku_blue.Println(`                        $$$"""""$$$                   $$$$"""$$$$$$$$"""""""""""""""""$$$$                               `)
	miku_blue.Println(`                        $$$"""""$$$                   $$$$""$$$$$$$$$$$$$""""""""""""""$$$                               `)
	miku_blue.Println(`                       $$$$"""""$$$$                  $$$"""$$$      $$$$$$$""""""""""""$$$                              `)
	miku_blue.Println(`                       $$$$""""""$$$                  $$$""$$$$          $$$$$;"""""""""$$$                              `)
	miku_blue.Print(`                   `);miku_pink.Print(`$$$$`);miku_blue.Println(`$$$$""""""$$$                  $$$""$$$             $$$$$""""""""$$$                              `)
	miku_blue.Print(`                `);miku_pink.Print(`$$$$$$$`);miku_blue.Println(`$$$$""""""$$$$                $$$$""$$$               $$$$"""""""$$$                              `)
	miku_blue.Print(`             `);miku_pink.Print(`$$$$$$$$$$`);miku_blue.Print(`$$$$"""""""$$$                $$$""$$$          `);miku_pink.Print(`$$`);miku_blue.Println(`     $$$$"""""$$$$                              `)
	miku_blue.Print(`           `);miku_pink.Print(`$$$$$$$$$$   `);miku_blue.Print(`$$$$"""""""$$$               $$$""$$$     `);miku_pink.Print(`$$$$$$$$`);miku_blue.Println(`     $$$""""$$$$                               `)
	miku_blue.Print(`         `);miku_pink.Print(`$$$$$$$$$$      `);miku_blue.Print(`$$$"""""""$$$$     $$$$$$   $$$"$$$$ `);miku_pink.Print(`$$$$$$$$$$$`);miku_blue.Println(`     $$$$"""$$$$                                `)
	miku_blue.Print(`       `);miku_pink.Print(`$$$$$j$$$$        `);miku_blue.Print(`$$$$"""""""$$$$$$$$$$$$$$$$$$$)"$$$`);miku_pink.Print(`$$$$$$$$$$$`);miku_blue.Println(`       $$$""$$$$$                                 `)
	miku_blue.Print(`      `);miku_pink.Print(`$$$$jj$$$$          `);miku_blue.Print(`$$$$""""""""$$$$$$""""""$$$$$""$$$`);miku_pink.Print(`$$$$$$$$`);miku_blue.Println(`         $$$$$$$$$                                   `)
	miku_blue.Print(`    `);miku_pink.Print(`$$$$$(($$$$            `);miku_blue.Print(`$$$$""""""$$$""""""""""""""""W$$$`);miku_pink.Print(`$$$$`);miku_blue.Println(`            $$$$$$$$                                     `)
	miku_blue.Print(`   `);miku_pink.Print(`$$$$j((($$$$$            `);miku_blue.Println(`$$$$$"""$$$"""""""""""""""""$$$$                $$$                                          `)
	miku_blue.Print(`   `);miku_pink.Print(`$$$j(((((j$$$$$$      $$$$$`);miku_blue.Println(`$$$$$$$$""""""""""""""""""$$$                                                              `)
	miku_blue.Print(`   `);miku_pink.Print(`$$$((((((((j$$$$$$$$$$$$$$$$$%`);miku_blue.Println(`$$$$W""""""""""""""""""$$$                                                              `)
	miku_blue.Print(`   `);miku_pink.Print(`$$$(((((((((((jjjjjjjjjj((((jj$`);miku_blue.Println(`$$$$""""""""""""""""""$$$                                                              `)
	miku_blue.Print(`   `);miku_pink.Print(`$$$$$j(((((((((((((((jj$$$$$$$$`);miku_blue.Println(`$$$$"""""""""""""""""a$$$$                  $$$$$                                      `)
	miku_blue.Print(`    `);miku_pink.Print(`$$$$$$$$$$$$$$$$$$$$$$$$$$$$`);miku_blue.Println(`   $$$$""""""""""""""""$$$$$$$$$$$$$$$$$$$$$$$$$$$$                                      `)
	miku_blue.Print(`       `);miku_pink.Print(`$$$$$$$$$$$$$$$$$$`);miku_blue.Println(`           $$$$$"""""""""""$$$$$"""z$$$$$$$$$$$$$$$$$$$$$                                       `)
	miku_blue.Println(`                                      $$$$$$$$$$$$$$$$$$$$z""""""""""z$$$$$$$$$                                          `)
	miku_blue.Println(`                                        $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$                                             `)
	miku_blue.Println(`                                                             $$$$$$$                                                     `)
	yellow.Println("\n欢迎使用 OllamaModeloader 模型加载器，当前版本：v1.0.0")
	yellow.Println("OllamaModeloader 是用于Ollama快速加载本地模型的工具，目前支持gguf格式的模型加载，以及创建模型文件。\n")
	yellow.Println("©create by @MakotoArai\n")
}

func checkOllama() {
	if _, err := exec.LookPath("ollama"); err != nil {
		red.Println("错误：ollama 未安装，请先安装ollama。")
		os.Exit(1)
	}
	green.Println("ollama 已安装。")
	yellow.Println("请确保当前目录包含以下内容：")
	yellow.Println("• 单文件的 .gguf 模型文件")
	yellow.Println("• ollama 支持的框架的模型文件夹\n")
}

func selectModel() string {
	files, _ := filepath.Glob("*.gguf")
	dirs, _ := filepath.Glob("*/")

	var items []string
	items = append(items, files...)
	for i, d := range dirs {
		dirs[i] = strings.TrimSuffix(d, "/")
	}
	items = append(items, dirs...)

	if len(items) == 0 {
		red.Println("错误：当前目录下未找到 .gguf 文件或模型文件夹。")
		os.Exit(1)
	}

	green.Println("\n可用的模型文件和文件夹：")
	for i, item := range items {
		fmt.Printf("%3d) %s\n", i+1, item)
	}
	red.Println("  q) 退出")

	for {
		fmt.Print("请选择要加载的模型编号或输入q退出: ")
		var choice string
		fmt.Scanln(&choice)

		if strings.ToLower(choice) == "q" {
			os.Exit(0)
		}

		if idx := parseInt(choice); idx > 0 && idx <= len(items) {
			return "./" + items[idx-1]
		}

		red.Printf("无效输入，请输入1-%d之间的数字或q退出\n", len(items))
	}
}

func modifyModelfile(path string) {
	content := fmt.Sprintf("FROM %s", path)
	if _, err := os.Stat("Modelfile"); err == nil {
		data, _ := os.ReadFile("Modelfile")
		re := regexp.MustCompile(`(?m)^FROM .*$`)
		content = re.ReplaceAllString(string(data), "FROM "+path)
		green.Println("已更新Modelfile中的模型路径")
	} else {
		green.Println("已创建新的Modelfile")
	}

	os.WriteFile("Modelfile", []byte(content), 0644)
}

func main() {
	checkOllama()
	selectedPath := selectModel()
	modifyModelfile(selectedPath)

	defaultName := strings.TrimSuffix(filepath.Base(selectedPath), ".gguf")
	green.Printf("请输入模型名称（建议使用英文，默认：%s）: ", defaultName)

	var modelName string
	fmt.Scanln(&modelName)
	if modelName == "" {
		modelName = defaultName
		yellow.Printf("使用默认名称: %s\n", modelName)
	}

	yellow.Println("\n正在创建模型...")
	exec.Command("ollama", "create", modelName, "-f", "./Modelfile").Run()

	green.Println("\n当前已安装模型列表：")
	exec.Command("ollama", "list").Run()
}

func parseInt(s string) int {
	var n int
	_, err := fmt.Sscanf(s, "%d", &n)
	if err != nil {
		return -1
	}
	return n
}
