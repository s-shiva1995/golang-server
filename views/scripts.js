function isValidName(name) {
  if (!name.trim().length) {
    document.getElementById("errorname").innerHTML = "Please enter name"
  } else {
    document.getElementById("errorname").innerHTML = ""
  }
}

function isValidPassword(password) {
  if (!password.trim().length) {
    document.getElementById("errorpassword").innerHTML = "Please enter password"
  } else {
    document.getElementById("errorpassword").innerHTML = ""
  }
}
