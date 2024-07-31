$('#login').on('submit', login);

function login(event) {
  event.preventDefault();

  $.ajax({
    url: '/login',
    method: 'POST',
    data: {
      email: $('#email').val(),
      passworld: $('#passworld').val(),
    },
  })
    .done(function () {
      window.location = '/home';
    })
    .fail(function () {
      alert('User or password invalid');
    });
}
