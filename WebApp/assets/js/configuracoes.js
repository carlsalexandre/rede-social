var USUARIO_LOGADO_ID = $('body').data('usuario-id');

$('#form-editar-perfil').on('submit', editarPerfil);
$('#form-alterar-senha').on('submit', alterarSenha);
$('#btn-excluir-conta').on('click', excluirConta);

function editarPerfil(evento) {
    evento.preventDefault();

    var form = $(this);

    $.ajax({
        url: "/usuarios/" + USUARIO_LOGADO_ID,
        method: "PUT",
        data: form.serialize()
    }).done(function() {
        Swal.fire({
            icon: 'success',
            title: 'Dados atualizados!',
            text: 'Suas informações foram salvas com sucesso.'
        }).then(function() {
            window.location.reload();
        });
    }).fail(function(erro) {
        console.log(erro);
        var mensagem = "Erro ao atualizar seus dados, tente novamente.";
        if (erro.responseJSON && erro.responseJSON.erro) {
            mensagem = erro.responseJSON.erro;
        }
        Swal.fire({
            icon: 'error',
            title: 'Não foi possível atualizar',
            text: mensagem
        });
    });
}

function alterarSenha(evento) {
    evento.preventDefault();

    var senhaNova = $('input[name="senhaNova"]').val();
    var confirmacao = $('input[name="senhaNovaConfirmacao"]').val();

    if (senhaNova !== confirmacao) {
        Swal.fire({
            icon: 'warning',
            title: 'Ops!',
            text: 'A confirmação de senha não bate com a nova senha.'
        });
        return;
    }

    var form = $(this);

    $.ajax({
        url: "/usuarios/" + USUARIO_LOGADO_ID + "/atualizar-senha",
        method: "POST",
        data: form.serialize()
    }).done(function() {
        Swal.fire({
            icon: 'success',
            title: 'Senha alterada!',
            text: 'Sua senha foi atualizada com sucesso.'
        });
        form[0].reset();
    }).fail(function(erro) {
        console.log(erro);
        var mensagem = "Erro ao alterar a senha, tente novamente.";
        if (erro.responseJSON && erro.responseJSON.erro) {
            mensagem = erro.responseJSON.erro;
        }
        Swal.fire({
            icon: 'error',
            title: 'Não foi possível alterar a senha',
            text: mensagem
        });
    });
}

function excluirConta() {
    Swal.fire({
        icon: 'warning',
        title: 'Excluir sua conta?',
        text: 'Essa ação não pode ser desfeita.',
        showCancelButton: true,
        confirmButtonText: 'Continuar',
        cancelButtonText: 'Cancelar',
        confirmButtonColor: '#e94560'
    }).then(function(primeiraConfirmacao) {
        if (!primeiraConfirmacao.isConfirmed) {
            return;
        }

        Swal.fire({
            icon: 'error',
            title: 'Tem certeza mesmo?',
            text: 'Sua conta e todas as suas publicações serão apagadas permanentemente.',
            showCancelButton: true,
            confirmButtonText: 'Sim, excluir minha conta',
            cancelButtonText: 'Cancelar',
            confirmButtonColor: '#e94560'
        }).then(function(segundaConfirmacao) {
            if (!segundaConfirmacao.isConfirmed) {
                return;
            }

            $.ajax({
                url: "/usuarios/" + USUARIO_LOGADO_ID,
                method: "DELETE"
            }).done(function() {
                Swal.fire({
                    icon: 'success',
                    title: 'Conta excluída',
                    text: 'Sentiremos sua falta!'
                }).then(function() {
                    window.location = "/login";
                });
            }).fail(function(erro) {
                console.log(erro);
                var mensagem = "Erro ao excluir sua conta, tente novamente.";
                if (erro.responseJSON && erro.responseJSON.erro) {
                    mensagem = erro.responseJSON.erro;
                }
                Swal.fire({
                    icon: 'error',
                    title: 'Não foi possível excluir',
                    text: mensagem
                });
            });
        });
    });
}