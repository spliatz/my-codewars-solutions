function shark(pontoonDistance, sharkDistance, youSpeed, sharkSpeed, dolphin) {
    let timeToPontoonPeople = pontoonDistance / youSpeed;
    let timeToPontoonShark = sharkDistance / sharkSpeed;
    let timeHalfSharkSpeed = sharkDistance / (sharkSpeed / 2);
    let timeDifferenceWithDolphin = timeHalfSharkSpeed - timeToPontoonPeople;
    let timeDifferenceWithoutDolphin = timeToPontoonShark - timeToPontoonPeople;

    if (sharkDistance < sharkSpeed) {
        return "Shark Bait!";
    } else if (pontoonDistance < youSpeed) {
        return "Alive!";
    }

    if (dolphin === true && timeDifferenceWithDolphin > 0) {
        return "Alive!";
    } else if (dolphin === true && timeDifferenceWithDolphin < 0) {
        return "Shark Bait!";
    } else if (dolphin === false && timeDifferenceWithoutDolphin > 0) {
        return "Alive!";
    } else {
        return "Shark Bait!";
    }
}