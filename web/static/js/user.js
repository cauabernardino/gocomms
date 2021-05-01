$('#unfollow').on('click', unfollow);
$('#follow').on('click', follow);

function unfollow() {
    const userID = $(this).data('user-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/users/${userID}/unfollow`,
        method: "POST"
    }).done(function () {
        window.location = `/users/${userID}`;
    }).fail(function () {
        Swal.fire({
            title: "Oops...",
            text: "Error unfollowing user!",
            icon: "error",
            confirmButtonColor: "#4e4e50",
        });
        $('#unfollow').prop('disabled', false)
    })
}

function follow() {
    const userID = $(this).data('user-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/users/${userID}/follow`,
        method: "POST"
    }).done(function () {
        window.location = `/users/${userID}`;
    }).fail(function () {
        Swal.fire({
            title: "Oops...",
            text: "Error following user!",
            icon: "error",
            confirmButtonColor: "#4e4e50",
        });
        $('#follow').prop('disabled', false)
    })

}
