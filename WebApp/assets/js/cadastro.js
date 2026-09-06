$('#cadastro').on('submit', criarUsuario)

function criarUsuario(evento) {
    evento.preventDefault();
    console.log("dentro da função usuario");

    if ($('#senha').val() != $('#confirmar-senha').val()) {
        alert("As senhas não estão iguais.");
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
        alert("Usuário cadastrado com sucesso, efetue o login");
    }).fail(function(erro) {
        console.log(erro);
        alert("Erro ao cadastrar usuário, verifique as informações");
    });
}