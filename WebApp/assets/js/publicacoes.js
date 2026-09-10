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