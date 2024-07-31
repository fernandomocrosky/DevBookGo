$('#register-form').on('submit', createUser);

function createUser(event) {
  event.preventDefault();
  console.log('Dentro da fução criar usuario');
  const password = $('#password').val();
  const confirmPassword = $('#confirm-password').val();

  if (password != confirmPassword) {
    alert("Passwords doesn't match!");
    return;
  }

  $.ajax({
    url: '/users',
    method: 'POST',
    data: {
      name: $('#name').val(),
      email: $('#email').val(),
      nick: $('#nick').val(),
      password: password,
    },
  })
    .done(function () {})
    .fail(function (err) {});
}
