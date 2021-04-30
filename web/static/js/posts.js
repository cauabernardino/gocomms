$('#new-post').on('submit', createPost);
$('.like-post').on('click', likePost);


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
    }).fail(function () {
        alert("Error on liking!")
    }).always(function () {
        clickedElement.prop('disabled', false);
    })
}