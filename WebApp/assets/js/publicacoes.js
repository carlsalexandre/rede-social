$('#nova-publicacao').on('submit', publicar);
 
function publicar(evento) {
    evento.preventDefault();
 
    var conteudo = $('#conteudo').val().trim();
 
    if (conteudo === '') {
        alert('Escreva algo antes de publicar.');
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
        alert("Erro ao criar a publicação, tente novamente.");
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
 
    if (!confirm('Tem certeza que deseja excluir esta publicação?')) {
        return;
    }
 
    $.ajax({
        url: "/publicacoes/" + publicacaoId,
        method: "DELETE"
    }).done(function() {
        window.location = "/home";
    }).fail(function(erro) {
        console.log(erro);
        alert("Erro ao excluir a publicação, tente novamente.");
    });
}