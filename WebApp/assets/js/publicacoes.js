$('#nova-publicacao').on('submit', publicar);

function publicar(evento) {
    evento.preventDefault();

    var conteudo = $('#conteudo').val().trim();

    if (conteudo === '') {
        Swal.fire({
            icon: 'warning',
            title: 'Ops!',
            text: 'Escreva algo antes de publicar.'
        });
        return;
    }

    $.ajax({
        url: "/publicacoes",
        method: "POST",
        data: {
            conteudo: conteudo
        }
    }).done(function() {
        window.location = "/home";
    }).fail(function(erro) {
        console.log(erro);
        Swal.fire({
            icon: 'error',
            title: 'Erro ao publicar',
            text: 'Tente novamente.'
        });
    });
}

$(document).on('click', '.app-post-menu-toggle', function (evento) {
    evento.stopPropagation();

    var painel = $(this).siblings('.app-post-menu-panel');
    var estaAberto = !painel.prop('hidden');

    $('.app-post-menu-panel').prop('hidden', true);
    $('.app-post-menu-toggle').attr('aria-expanded', 'false');

    if (!estaAberto) {
        painel.prop('hidden', false);
        $(this).attr('aria-expanded', 'true');
    }
});

$(document).on('click', function () {
    $('.app-post-menu-panel').prop('hidden', true);
    $('.app-post-menu-toggle').attr('aria-expanded', 'false');
});

$(document).on('click', '.app-post-menu-panel', function (evento) {
    evento.stopPropagation();
});

$(document).on('click', '.app-post-menu-delete', excluirPublicacao);

function excluirPublicacao(evento) {
    evento.preventDefault();

    var publicacaoId = $(this).data('publicacao-id');

    Swal.fire({
        icon: 'warning',
        title: 'Excluir publicação?',
        text: 'Essa ação não pode ser desfeita.',
        showCancelButton: true,
        confirmButtonText: 'Sim, excluir',
        cancelButtonText: 'Cancelar',
        confirmButtonColor: '#e94560'
    }).then(function(resultado) {
        if (!resultado.isConfirmed) {
            return;
        }

        $.ajax({
            url: "/publicacoes/" + publicacaoId,
            method: "DELETE"
        }).done(function() {
            window.location = "/home";
        }).fail(function(erro) {
            console.log(erro);
            Swal.fire({
                icon: 'error',
                title: 'Erro ao excluir',
                text: 'Tente novamente.'
            });
        });
    });
}

$(document).on('click', '.curtir-publicacao', curtirOuDescurtir);

function curtirOuDescurtir(evento) {
    evento.preventDefault();

    var botao = $(this);

    if (botao.prop('disabled')) {
        return;
    }

    var publicacaoId = botao.data('publicacao-id');
    var jaCurtida = botao.hasClass('is-curtido');
    var acao = jaCurtida ? 'descurtir' : 'curtir';
    var contador = botao.find('span');
    var curtidasAtuais = parseInt(contador.text(), 10) || 0;

    botao.prop('disabled', true);

    $.ajax({
        url: "/publicacoes/" + publicacaoId + "/" + acao,
        method: "POST"
    }).done(function() {
        botao.toggleClass('is-curtido');
        contador.text(jaCurtida ? curtidasAtuais - 1 : curtidasAtuais + 1);
    }).fail(function(erro) {
        console.log(erro);
        Swal.fire({
            icon: 'error',
            title: 'Erro ao curtir',
            text: 'Tente novamente.'
        });
    }).always(function() {
        botao.prop('disabled', false);
    });
}