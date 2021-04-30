$('#new-post').on('submit', createPost);
$(document).on('click', '.like-post', likePost)
$(document).on('click', '.unlike-post', unlikePost)
$('#send-edit').on('click', editPost);
$('.delete-post').on('click', deletePost)

function createPost(event) {
    event.preventDefault();


    $.ajax({
        url: "/posts",
        method: "POST",
        data: {
            "title": $('#title').val(),
            "content": $('#content').val(),
        }
    }).done(function () {
        window.location = "/home";
    }).fail(function () {
        Swal.fire({
            title: "Oops...",
            text: "Error on creating the post!",
            icon: "error",
            confirmButtonColor: "#4e4e50",
        });
    });
}

function likePost(event) {
    event.preventDefault();

    const clickedElement = $(event.target);
    const postID = clickedElement.closest('div').data('post-id');

    clickedElement.prop('disabled', true);

    $.ajax({
        url: `/posts/${postID}/like`,
        method: "POST"
    }).done(function () {
        const countLike = clickedElement.next('span');
        const qtyLike = parseInt(countLike.text());

        countLike.text(qtyLike + 1);

        // For unliking
        clickedElement.addClass('unlike-post');
        clickedElement.addClass('like-color');
        clickedElement.removeClass('like-post');

    }).fail(function () {
        Swal.fire({
            title: "Oops...",
            text: "Error on the post!",
            icon: "error",
            confirmButtonColor: "#4e4e50",
        });
    }).always(function () {
        clickedElement.prop('disabled', false);
    })
}

function unlikePost(event) {
    event.preventDefault();

    const clickedElement = $(event.target);
    const postID = clickedElement.closest('div').data('post-id');

    clickedElement.prop('disabled', true);

    $.ajax({
        url: `/posts/${postID}/unlike`,
        method: "POST"
    }).done(function () {
        const countLike = clickedElement.next('span');
        const qtyLike = parseInt(countLike.text());

        countLike.text(qtyLike - 1);

        // For unlinking
        clickedElement.removeClass('unlike-post');
        clickedElement.removeClass('like-color');
        clickedElement.addClass('like-post');

    }).fail(function () {
        Swal.fire({
            title: "Oops...",
            text: "Error on unliking the post!",
            icon: "error",
            confirmButtonColor: "#4e4e50",
        })
    }).always(function () {
        clickedElement.prop('disabled', false);
    })
}

function editPost(event) {
    $(this).prop('disabled', true);

    const postID = $(this).data('post-id')

    $.ajax({
        url: `/posts/${postID}`,
        method: "PUT",
        data: {
            title: $('#title').val(),
            content: $('#content').val(),
        }
    }).done(function () {
        Swal.fire({
            title: "Success!",
            text: "Post edited successfully!",
            icon: "success",
            confirmButtonColor: "#4e4e50",
        }).then(function () {
            window.location = "/home"
        });
    }).fail(function () {
        Swal.fire({
            title: "Oops...",
            text: "Error on editing the post!",
            icon: "error",
            confirmButtonColor: "#4e4e50",
        })
    }).always(function () {
        $("#send-edit").prop('disabled', false)
    })

}

function deletePost(event) {
    event.preventDefault();

    Swal.fire({
        title: "Are you sure?",
        text: "After deleting, you won't be able to see this message again.",
        icon: "warning",
        showCancelButton: true,
        cancelButtonColor: '#950740',
        cancelButtonText: "Cancel",
        confirmButtonColor: "#4e4e50",
    }).then(function (confirmation) {
        if (!confirmation.value) return;

        const clickedElement = $(event.target);
        const post = clickedElement.closest('div')
        const postID = post.data('post-id');

        clickedElement.prop('disabled', true);

        $.ajax({
            url: `/posts/${postID}`,
            method: "DELETE"
        }).done(function () {
            post.fadeOut("slow", function () {
                $(this).remove();
            });
        }).fail(function () {
            Swal.fire({
                title: "Oops...",
                text: "Error on deleting post!",
                icon: "error",
                confirmButtonColor: "#4e4e50",
            });
        });
    })

}