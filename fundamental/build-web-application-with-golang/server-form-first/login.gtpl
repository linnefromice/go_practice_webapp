<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Login</title>
</head>
<body>
    <form action="/login" method="post">
        <input type="hidden" name="token" value ="{{.}}">
        <p>ユーザ名: <input type="text" name="username"></p>
        <p>パスワード: <input type="password" name="password"></p>
        <p><input type="submit" value="ログイン"></p>
    </form>
</body>
</html>