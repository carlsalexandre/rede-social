$('#cadastro').on('submit', criarUsuario)

function criarUsuario(evento) {
    evento.preventDefault();

    if ($('#senha').val() != $('#confirmar-senha').val()) {
        Swal.fire({
            icon: 'warning',
            title: 'Ops!',
            text: 'As senhas não estão iguais.'
        });
        return;
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            nome:   $('#nome').val(),
            email:  $('#email').val(),
            nick:   $('#nick').val(),
            senha:  $('#senha').val()
        }
    }).done(function() {
        Swal.fire({
            icon: 'success',
            title: 'Conta criada!',
            text: 'Usuário cadastrado com sucesso, efetue o login.'
        }).then(function() {
            window.location = "/login";
        });
    }).fail(function(erro) {
        console.log(erro);
        Swal.fire({
            icon: 'error',
            title: 'Erro ao cadastrar',
            text: 'Verifique as informações e tente novamente.'
        });
    });
}