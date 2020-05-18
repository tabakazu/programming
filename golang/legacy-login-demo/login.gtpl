<!DOCTYPE html>
<html>
		<h1>Login</h1>
		<form action="/login" method="post">
	      ユーザ名:<input type="text" name="username">
        パスワード:<input type="password" name="password">
        <input type="hidden" name="token" value="{{.}}">
	      <input type="submit" value="ログイン">
    </form>
</html>

