# Go-Miiko

[![pipeline status](https://gitlab.com/NatoBoram/Go-Miiko/badges/master/pipeline.svg)](https://gitlab.com/NatoBoram/Go-Miiko/commits/master)
[![Go Report Card](https://goreportcard.com/badge/gitlab.com/NatoBoram/Go-Miiko)](https://goreportcard.com/report/gitlab.com/NatoBoram/Go-Miiko)
[![GoDoc](https://godoc.org/gitlab.com/NatoBoram/Go-Miiko?status.svg)](https://godoc.org/gitlab.com/NatoBoram/Go-Miiko)

Go-Miiko is a [Discord](https://discordapp.com/) bot that takes care of [Eldarya](http://www.eldarya.fr/) themed Discord servers. It's written in [Go](https://golang.org/) with the help of [DiscordGo](https://github.com/bwmarrin/discordgo). Right now, she only speaks French. As a result, she can only be used in a French server.

You can invite her by clicking [here](https://discordapp.com/api/oauth2/authorize?client_id=376971915010768896&permissions=268946499&scope=bot). She might come off offline fairly often as I'm working on her and my dedicated server isn't stable.

I'm working on bringing her to a state where she can create and manage a Discord Server.

![Miiko](assets/Miiko.png)

## Features

* Welcome newcomers
* Place newcomers in their guard
* Create said guard if it doesn't exist
* Warn Server Owner when there's a Light Guard member
* Pin popular messages
* Send received direct messages to its master
* Send an invite to people who leave the server
* Tells people when she's typing
* Likes popcorn
* ~~Likes candies~~
* Has *lots* of quotes. Like, **a lot**.

## Installation

Go-Miiko works best with the followings :

* MariaDB
* NginX
* PHPMyAdmin

The installation instructions are based on Digital Ocean's guides.

1. [How To Install Linux, Nginx, MySQL, PHP (LEMP stack) on Ubuntu 18.04](https://www.digitalocean.com/community/tutorials/how-to-install-linux-nginx-mysql-php-lemp-stack-ubuntu-18-04)
2. [How To Install and Secure phpMyAdmin with Nginx on Ubuntu 16.04](https://www.digitalocean.com/community/tutorials/how-to-install-and-secure-phpmyadmin-with-nginx-on-ubuntu-16-04)

Note that some changes needs to be made to these guides to use MariaDB instead of MySQL.

### MariaDB

```bash
sudo apt install nginx
sudo apt install mariadb-server
sudo mysql_secure_installation
```

Now that MariaDB is installed, you need to set-up a user because `root` can't always be used.

```bash
sudo mysql
```

You can check the users on your database with the following command :

```sql
SELECT user, authentication_string, plugin, host FROM mysql.user;
```

```sql
+------+-----------------------+-------------+-----------+
| user | authentication_string | plugin      | host      |
+------+-----------------------+-------------+-----------+
| root |                       | unix_socket | localhost |
+------+-----------------------+-------------+-----------+
1 row in set (0.00 sec)
```

Now, let's create a user and give it some permissions.

```sql
CREATE USER 'enter_user_here'@'%' IDENTIFIED BY 'enter_password_here';
GRANT ALL PRIVILEGES ON *.* TO 'enter_user_here'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES;
SELECT user, authentication_string, plugin, host FROM mysql.user;
```

```sql
+------+-----------------------+-------------+-----------+
| user | authentication_string | plugin      | host      |
+------+-----------------------+-------------+-----------+
| root |                       | unix_socket | localhost |
| nato |                       |             | %         |
+------+-----------------------+-------------+-----------+
2 rows in set (0.00 sec)
```

You can replace `%` by `localhost` if you want your new root user to connect only from localhost instead of from everywhere. This is useful in testing environments, but you might need access from outside if you're installing the bot elsewhere.

### Nginx

These steps differ from one installation to another.

* If you have a domain name, then please follow the [guide](https://www.digitalocean.com/community/tutorials/how-to-install-linux-nginx-mysql-php-lemp-stack-ubuntu-18-04#step-3-%E2%80%93-installing-php-and-configuring-nginx-to-use-the-php-processor).
* If you plan on getting a domain name but want to set-up your stuff before acquiring it, then let me tell you this : It's a bad idea, stop right now, and get your domain name, then follow the guide above.
* If this machine will never get a domain name (it's a dev machine), then follow these instructions.

```bash
sudo apt install php-fpm php-mysql
```

We're going to edit the `default` site so that it always use PHP no matter the website.

```bash
sudo nano /etc/nginx/sites-available/default
```

Let's tell Nginx to use PHP.

```nginx
	# Add index.php to the list if you are using PHP
	index index.php index.html index.htm index.nginx-debian.html;
```

Now it will try to use it, but it won't know what to do with it. You have to enable it!

```nginx
	# pass PHP scripts to FastCGI server

	location ~ \.php$ {
		include snippets/fastcgi-php.conf;

		# With php-fpm (or other unix sockets):
		fastcgi_pass unix:/var/run/php/php7.0-fpm.sock;
		# With php-cgi (or other tcp sockets):
		#fastcgi_pass 127.0.0.1:9000;
	}
```

Now, notice the `php7.0-fpm.sock`. You are *probably* using a different version, so you **have** to change it. Let's see how it's done.

```bash
apt show php-fpm
```

```yaml
Depends: php7.2-fpm
```

Replace `php7.0-fpm.sock` with your `php-fpm`. It my case, it's `php7.2-fpm.sock`.

This part is uncommented generally when you have Apache installed next to Nginx. It is a good practice to *always* uncomment it, just in case.

```nginx
	# deny access to .htaccess files, if Apache's document root
	# concurs with nginx's one

	location ~ /\.ht {
	       deny all;
	}
```

Now that these changes were made, it's time to apply them. First, test them.

```bash
sudo nginx -t
```

```log
nginx: the configuration file /etc/nginx/nginx.conf syntax is ok
nginx: configuration file /etc/nginx/nginx.conf test is successful
```

If there's an error, then you've done something wrong. Check your configuration.

Now it's time to restart Nginx.

```bash
sudo systemctl reload nginx
```

You can test your configuration by [Creating a PHP File to Test Configuration](https://www.digitalocean.com/community/tutorials/how-to-install-linux-nginx-mysql-php-lemp-stack-ubuntu-18-04#step-4-%E2%80%93-creating-a-php-file-to-test-configuration).

### PHPMyAdmin

```bash
sudo apt install phpmyadmin
```

The installation will ask you to choose the web server that should be automatically configured to run phpMyAdmin. Since we're using Nginx and it's not in the choices, press <kbd>Tab</kbd> then <kbd>Enter</kbd>.

Accept the database configuration, and leave the password blank so it generates a secure password unique to PHPMyAdmin. Since you have your own user, you don't need to login as PHPMyAdmin, so setting a weaker password is useless.

Now let's add PHPMyAdmin to your Nginx default website.

```bash
sudo ln -s /usr/share/phpmyadmin /var/www/html
```

Apparently, PHPMyAdmin uses `mycrypt`, but I don't really know if it's true. Nontheless, let's run the commands we're supposed to run, just in case.

```bash
sudo apt install mcrypt
sudo phpenmod mcrypt
sudo systemctl restart php7.2-fpm
```

Use the `php7.2-fpm` version we discovered earlier with `apt show php-fpm`.

You can enhance your server's security by [Change the Default phpMyAdmin URL](https://www.digitalocean.com/community/tutorials/how-to-install-and-secure-phpmyadmin-with-nginx-on-ubuntu-16-04#step-2-%E2%80%94-change-the-default-phpmyadmin-url) and by [Set Up an Nginx Authentication Gateway](https://www.digitalocean.com/community/tutorials/how-to-install-and-secure-phpmyadmin-with-nginx-on-ubuntu-16-04#step-3-%E2%80%94-set-up-an-nginx-authentication-gateway), but since this guide only covers the localhost for developement purpose, this will be all for now.

### Setting up Go-Miiko

Go-Miiko needs a database, so let's create a user in the server.

* User name : Miiko
* Host name : localhost
* Password : 2y7KZX4wqhEJ3O4J
* Re-type: 2y7KZX4wqhEJ3O4J
* Authentication Plugin : Native MySQL authentication
* Generate password : 2y7KZX4wqhEJ3O4J
* [x] Create database with same name and grant all privileges.

Please generate a random password and note it somewhere. Since you won't log-in as Miiko, it's a good practice to use a complex randomly generated password instead of entering a weaker one.

Go-Miiko doesn't set-up her database by herself at the moment. You need to run `Miiko.sql` manually.

Now that its user is done, we need to enter its configuration where the bot should run. Ideally, it'll be in `~/Miiko/`.

First, install the bot and run it once where you want its files located.

```bash
cd
go get -u -v -fix gitlab.com/NatoBoram/Go-Miiko
Go-Miiko
```

```log
Could not load the database configuration.
open ./Miiko/db.json: no such file or directory
Writing a new database configuration template...
```

The bot created a template where to insert its configuration. Add the necessary info there.

```bash
nano Miiko/db.json
```

```json
{
	"User": "Miiko",
	"Password": "2y7KZX4wqhEJ3O4J",
	"Address": "localhost",
	"Port": 3306,
	"Database": "Miiko"
}
```

Once the database is correctly configured, run the bot again to check if it's working.

```bash
Go-Miiko
```

```log
Could not load the Discord configuration.
open ./Miiko/discord.json: no such file or directory
Writing a new Discord configuration template...
```

This means the database was used successfully, and it is now asking for its Discord configuration. Get your token [here](https://discordapp.com/developers/applications/) and your User ID [here](https://support.discordapp.com/hc/en-us/articles/206346498-Where-can-I-find-my-User-Server-Message-ID-).

```json
{
	"Token": "NDAwMDU2NzIzOTAzMDg2NTky.DoVKGA.MEV8ip4iM_8DL3lpINpxNML0jTs",
	"MasterID": "157311679288311808"
}
```

Then again, run the bot to check if everything is okay.

```bash
Go-Miiko
```

> Hi, master NatoBoram. I am Miiko, and everything's all right!

Congratulations! This means Go-Miiko is connected to Internet and is working as intended!

You can now start development, or invite it to your own server.