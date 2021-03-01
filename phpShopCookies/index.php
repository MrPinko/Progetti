<?php 
    $name = htmlspecialchars($_COOKIE['nome']);
    if(isset($name))
        header("location:shop.php");



?>

<html>
<head>
    <h1><center>LOGIN</center> </h1>
</head>

<body>


<form method="post" action="<?php echo $_SERVER['PHP_SELF']; ?>">
    nome: <input type="text" id="name" name="name"><br><br>
    cognome: <input type="text" id="surname" name="surname"><br><br>
    <p>sesso:</p>

            <input type="radio" id="genderBox" name="genderBox" value="M">
            <label>Maschio</label>
            <br>
            <input type="radio" id="genderBox" name="genderBox" value="F">
            <label>Femmina</label>
            <br>

    <br>
    <hr>

    <h3> HOBBY </h3>
            <input type="checkbox" name="hobby[0]" value="1">
            <label>Cucina</label>
            <br>
            <input type="checkbox" name="hobby[1]" value="1">
            <label>informatica</label>
            <br>
            <input type="checkbox" name="hobby[2]" value="1">
            <label>barca a vela</label>
            <br>
            <input type="checkbox" name="hobby[3]" value="1">
            <label>case</label>

    <br><br>
    <input type="submit" >
</form>



<?php
if ($_SERVER["REQUEST_METHOD"] == "POST") {
  // collect value of input field
    $name = $_REQUEST['name'];
    $surname = $_REQUEST['surname'];
    $gender = $_REQUEST['genderBox'];
    $hobbyList = array(
        'Cucina' => isset($_REQUEST['hobby'][0]),
        'Informatica' => isset($_REQUEST['hobby'][1]),
        'barca a vela' => isset($_REQUEST['hobby'][2]),
        'case' => isset($_REQUEST['hobby'][3]),
    );
    
    /*
    echo $name;
    echo '<br>';
    echo $surname;
    echo '<br>';
    echo $gender;
    echo '<br>';
    echo json_encode($hobbyList);
    setcookie('nome', $name);
    setcookie('cognome', $surname);
    setcookie('genderBox', $gender);
    setcookie('hobby', json_encode($hobbyList));
    */

    header("location:shop.php");

}

?>


</body>

</html>