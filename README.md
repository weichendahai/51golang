#sqlite3 基本使用
## 链接数据库
sqlite3 testDB.db
##退出 sqlite
sqlite>.quit

##命令来检查它是否在数据库列表中
sqlite>.databases

##附加数据库
sqlite> ATTACH DATABASE 'testDB.db' as 'testDB';

##分离数据库
sqlite> DETACH DATABASE 'currentDB';

##查看所有表
sqlite> .tables

##格式化输出 
sqlite>.header on
sqlite>.mode column
sqlite>.timer on

##显示各种设置的当前值
sqlite>.show

# golang jin 测试

    curl localhost:8005/simple/server/get
    curl localhost:8005/simple/server/get?key=hello
    curl -d "" localhost:8005/simple/server/post
    curl --request PUT localhost:8005/simple/server/put
    curl --request DELETE localhost:8005/simple/server/delete

#git

    …or create a new repository on the command line

    echo "# 51golang" >> README.md
    git init
    git add README.md
    git commit -m "first commit"
    git remote add origin https://github.com/weichendahai/51golang.git
    git push -u origin master
    …or push an existing repository from the command line

    git remote add origin https://github.com/weichendahai/51golang.git
    git push -u origin master

#FAQ
1.包内方法大小写错误；导致引用包方法失败
weichen@172-15-70-14 ~/src/go_workspace/51golang (master*) $ sh a.sh
 command-line-arguments
./main.go:80: cannot refer to unexported name models.getAllArticles
./main.go:80: undefined: models.getAllArticles

