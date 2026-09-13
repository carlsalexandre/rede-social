$('#login').on('submit', fazerlogin);

function fazerlogin(evento) {
    evento.preventDefault();

    $.ajax({
        url: "/login",
       method: "POST",
       data: {
            email: $('#email').val(),
            senha: $('#senha').val(),
       }
    }).done(function() {
        window.location = "/home"
    }).fail(function() {
        Swal.fire({
            icon: 'error',
            title: 'Não foi possível entrar',
            text: 'Usuário ou senha inválido.'
        });
    })
}