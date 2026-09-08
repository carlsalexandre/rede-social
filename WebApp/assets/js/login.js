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
        alert("Usuário/Senha inválido.")
    })
}