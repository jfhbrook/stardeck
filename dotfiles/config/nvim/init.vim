set backspace=indent,eol,start

" vim-plug
call plug#begin(stdpath('data') . '/plugged')
  Plug 'junegunn/fzf', { 'do': { -> fzf#install() } }
  Plug 'junegunn/fzf.vim'
  Plug 'liuchengxu/vim-which-key'
  Plug 'mileszs/ack.vim'
  Plug 'neoclide/coc.nvim', { 'branch': 'release' }
  Plug 'timonv/vim-cargo'
  Plug 'vmchale/just-vim'
  Plug 'Mofiqul/dracula.nvim'
  Plug 'hashivim/vim-terraform'
  Plug 'ray-x/go.nvim'
  Plug 'onyx-lang/onyx.vim'
  Plug 'lepture/vim-jinja'
call plug#end()

" force some filetypes
autocmd BufNewFile,BufReadPost *.md set filetype=markdown

" tab settings
set expandtab
set shiftwidth=2
set tabstop=2
set softtabstop=2

" gotta have dem line numbers
set nu
set ruler

" I hate swap files
set noswapfile

" des colores...
syntax on
set termguicolors
set colorcolumn=80
colorscheme dracula

" ag integration
if executable('ag')
  let g:ackprg = 'ag --vimgrep'
endif

" onyx language
augroup onyx_ft
  au!
  autocmd BufNewFile,BufRead *.onyx set syntax=onyx
augroup END

" jinja/nunjucks
au BufNewFile,BufRead *.j2, *.njk set ft=jinja

" coc.nvim

" an incantation
set signcolumn=number

" trigger coc.nvim autocomplete w/ this incantantion
function! CheckBackspace() abort
  let col = col('.') - 1
  return !col || getline('.')[col - 1] =~# '\s'
endfunction

inoremap <silent><expr> <Tab>
      \ coc#pum#visible() ? coc#pum#next(1) :
      \ CheckBackspace() ? "\<Tab>" :
      \ coc#refresh()

" vim-which-key

" This block is extremely brittle, leave it alone
call which_key#register('<Space>', "g:which_key_map")
let g:mapleader = "\<Space>"
let g:maplocalleader = ','
nnoremap <silent> <leader> :<c-u>WhichKey '<Space>'<CR>
nnoremap <silent> <localleader> :<c-u>WhichKey ','<CR>
" ok now we're good

" which-key bindings
let g:which_key_map = {}
let g:which_key_map['b'] = {
      \ 'name': '+buffers',
      \ 'd': ['bd', 'delete buffer'],
      \ 'l': [':BLines', 'fzf lines (this buffer)'],
      \ 'L': [':Lines', 'fzf lines (all buffers)'],
      \ '?': [':Buffers', 'fzf buffers'],
      \ }
let g:which_key_map['e'] = {
      \ 'name': '+exec',
      \ }
let g:which_key_map['f'] = {
      \ 'name': '+files',
      \ 's': 'save-file',
      \ '?': [':Files', 'fzf files']
      \ }
let g:which_key_map['w'] = {
      \ 'name' : '+windows' ,
      \ 'w' : ['<C-W>w'     , 'other-window']          ,
      \ 'd' : ['<C-W>c'     , 'delete-window']         ,
      \ '-' : ['<C-W>s'     , 'split-window-below']    ,
      \ '|' : ['<C-W>v'     , 'split-window-right']    ,
      \ '2' : ['<C-W>v'     , 'layout-double-columns'] ,
      \ 'h' : ['<C-W>h'     , 'window-left']           ,
      \ 'j' : ['<C-W>j'     , 'window-below']          ,
      \ 'l' : ['<C-W>l'     , 'window-right']          ,
      \ 'k' : ['<C-W>k'     , 'window-up']             ,
      \ 'H' : ['<C-W>5<'    , 'expand-window-left']    ,
      \ 'J' : [':resize +5'  , 'expand-window-below']   ,
      \ 'L' : ['<C-W>5>'    , 'expand-window-right']   ,
      \ 'K' : [':resize -5'  , 'expand-window-up']      ,
      \ '=' : ['<C-W>='     , 'balance-window']        ,
      \ 's' : ['<C-W>s'     , 'split-window-below']    ,
      \ 'v' : ['<C-W>v'     , 'split-window-below']    ,
      \ '?' : ['Windows'    , 'fzf-window']            ,
      \ }
