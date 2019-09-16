# Nagus
A golang app for updating a local git repo via http request


## installation setup
* create a user 'nagus' with home directory in /var/lib/nagus
* give user an ssh key pair
* put user's public key in github repo
* have github hit url in a hook:  my.ip.add.ress:8080/build?user=username&repo=reponame

## hook example
* http://nagus.cartcollector.com:8080/build?user=picardrulez&repo=kalinske&gobuild=true
