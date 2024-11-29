class AsyncQueue {
    #queue = [];
    #runningJob = [];
    #interval;
    #isProcessing = false;

    #stop = false;

    /**
     * @param {number} concurentNumber 
     */
    constructor(concurentNumber) {
    }

    /**
     * @param {() => void} task 
     */
    add(task) {
        this.#queue.push(task);

        if (!this.#isProcessing) {
            this.processQueue();
        }
    }

    async processQueue() {
        console.log('start processing queue')
        while (this.#queue.length > 0) {
            console.log('continue processing')
            const task = this.#queue.shift();

            try {
                await task();
            } catch (e) {
                console.log(e);
            }
        }
        console.log('stop processing queue')
        this.#isProcessing = false;
    }

    // start() {
    //     if (!this.#isProcessing) return;

    //     this.#isProcessing = true;
    //     this.processQueue();
    // }

    stop() {
        this.#isProcessing = false;
    }
}

function main() {
    const jobNum = 10;
    const aq = new AsyncQueue(4);

    for (let i = 1; i <= jobNum; i++) {
        aq.add(() => fetchMock(i));
    }

    setTimeout(() => aq.stop(), 10000);
}

/**
 * @param {number} id 
 */
async function fetchMock(id) {
    await new Promise((resolve) => setTimeout(resolve, 1000));

    await 
        fetch(`https://jsonplaceholder.typicode.com/todos/${id}`)
        .then(response => response.json())
        .then(json => console.log(json))
}

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

main();