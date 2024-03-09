function jungle(param){
    let currentSpecies = [];
    let predator = [-1];
    let count = 0;

    while (true) {
        for (let i = 0; i < param.length; i++) {
            for (let j = 0; j < predator.length; j++) {
                if (param[i] === predator[j]) {
                    currentSpecies.push(i);
                }
            }
        }

        if (currentSpecies.length > 0) {
            predator = [...currentSpecies]
            currentSpecies = [];
            count++;
        } else {
            break;
        }
    }

    console.log(count);
}

jungle([-1, 8, 6, 0, 7, 3, 8, 9, -1, 6]);
//       0  1  2  3  4  5  6  7   8  9
// 0, 8

// 3



