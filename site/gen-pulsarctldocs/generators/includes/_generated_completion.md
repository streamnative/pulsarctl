
------------

# completion





### Usage

`$ completion`



------------

## <em>bash</em>





 To load completion run 

  

 . <(pulsarctl completion bash) 

  

 To configure your bash shell to load completions for each session add to your bashrc 

  
 // ~/.bashrc or ~/.profile 

 . <(pulsarctl completion bash) 

  

 If you are stuck on Bash 3 (macOS) use 

  

 source /dev/stdin <<<"$(pulsarctl completion bash)" 

 

### Usage

`$ bash`




------------

## <em>zsh</em>





 To configure your zsh shell, run: 

  

 mkdir -p ~/.zsh/completion/ 

 pulsarctl completion zsh > ~/.zsh/completion/_pulsarctl 

  

 and put the following in ~/.zshrc: 

  

 fpath=($fpath ~/.zsh/completion) 

 

### Usage

`$ zsh`





