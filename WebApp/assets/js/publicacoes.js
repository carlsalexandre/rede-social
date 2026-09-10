$('#nova-publicacao').on('submit', publicar);
 
function publicar(evento) {
    evento.preventDefault();
 
    var titulo = $('#titulo').val().trim();
    var conteudo = $('#conteudo').val().trim();
 
    if (titulo === '' || conteudo === '') {
        alert('Preencha o título e o conteúdo da publicação.');
        return;
    }
 
    $.ajax({
        url: "/publicacoes",
        method: "POST",
        data: {
            titulo: titulo,
            conteudo: conteudo
        }
    }).done(function() {
        window.location = "/home";
    }).fail(function(erro) {
        console.log(erro);
        alert("Erro ao criar a publicação, verifique as informações e tente novamente.");
    });
}