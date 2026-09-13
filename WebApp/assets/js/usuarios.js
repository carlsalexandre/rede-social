$(document).on('click', '.seguir-usuario', seguirOuDeixarDeSeguir);

function seguirOuDeixarDeSeguir(evento) {
    evento.preventDefault();

    var botao = $(this);

    if (botao.prop('disabled')) {
        return;
    }

    var usuarioId = botao.data('usuario-id');
    var jaSegue = botao.hasClass('is-seguindo');
    var acao = jaSegue ? 'parar-de-seguir' : 'seguir';

    botao.prop('disabled', true);

    $.ajax({
        url: "/usuarios/" + usuarioId + "/" + acao,
        method: "POST"
    }).done(function() {
        window.location.reload();
    }).fail(function(erro) {
        console.log(erro);
        Swal.fire({
            icon: 'error',
            title: 'Erro ao atualizar',
            text: 'Tente novamente.'
        });
        botao.prop('disabled', false);
    });
}