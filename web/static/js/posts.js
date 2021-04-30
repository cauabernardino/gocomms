$('#new-post').on('submit', createPost);
$(document).on('click', '.like-post', likePost)
$(document).on('click', '.unlike-post', unlikePost)
$('#send-edit').on('click', editPost);


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
        alert("Error on creating the post!");
    })
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
        alert("Error on liking!")
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
        alert("Error on liking!")
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
        alert("Post edited successfully!");
        window.location = "/home";
    }).fail(function () {
        alert("Error on editing post!");
    }).always(function () {
        $("#send-edit").prop('disabled', false)
    })

}