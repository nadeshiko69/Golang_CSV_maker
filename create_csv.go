package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"slices"
	"strconv"
	"sync"
	"time"
)

func create_ext_code(wg *sync.WaitGroup, ch chan<- []string) {
	defer wg.Done()
	defer close(ch)

	var ext_code_list = []string{}
	var ext_code_for_output = []string{}

	// EXT_CODEの候補を必要数作成
	ext_cnt := 0
	for ext_cnt < EXT_CODE_COUNT {
		num := rand.Intn(999999)
		if slices.Contains(ext_code_list, fmt.Sprintf("%06d", num)) == false {
			ext_code_list = append(ext_code_list, fmt.Sprintf("%06d", num))
			ext_cnt++
		}
	}

	// 出力用のリスト作成
	list_cnt := 0
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for list_cnt < ROW_COUNT {
		random_index := r.Intn(EXT_CODE_COUNT)
		ext_code_for_output = append(ext_code_for_output, ext_code_list[random_index])
		list_cnt++
	}

	ch <- ext_code_for_output
	fmt.Println("create_ext_code Done")
}

func create_b_code(wg *sync.WaitGroup, ch chan<- []string) {
	defer wg.Done()
	defer close(ch)
	var b_code_list = []string{}
	var b_code_for_output = []string{}

	// EXT_CODEの候補を必要数作成
	b_code_list = append(b_code_list, "0123")
	b_cnt := 0
	for b_cnt < B_CODE_COUNT-1 {
		num := rand.Intn(300)
		if slices.Contains(b_code_list, fmt.Sprintf("%04d", num)) == false {
			b_code_list = append(b_code_list, fmt.Sprintf("%04d", num))
			b_cnt++
		}
	}

	// 出力用のリスト作成
	list_cnt := 0
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for list_cnt < ROW_COUNT {
		random_index := r.Intn(B_CODE_COUNT)
		b_code_for_output = append(b_code_for_output, b_code_list[random_index])
		list_cnt++
	}

	ch <- b_code_for_output

	fmt.Println("create_b_code Done")
}

func create_d_type(wg *sync.WaitGroup, ch chan<- []int) {
	defer wg.Done()
	defer close(ch)
	var d_type_list = []int{}
	var d_type_for_output = []int{}

	min := 10
	max := 40

	d_cnt := 0
	for d_cnt < D_TYPE_COUNT {
		num := min + rand.Intn(max+min-1)
		if slices.Contains(d_type_list, num) == false {
			d_type_list = append(d_type_list, num)
			d_cnt++
		}
	}

	// 出力用のリスト作成
	list_cnt := 0
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for list_cnt < ROW_COUNT {
		random_index := r.Intn(D_TYPE_COUNT)
		d_type_for_output = append(d_type_for_output, d_type_list[random_index])
		list_cnt++
	}

	ch <- d_type_for_output

	fmt.Println("create_d_type Done")
}

func create_n_time(wg *sync.WaitGroup, ch chan<- []int) {
	defer wg.Done()
	defer close(ch)
	var n_time_list = []int{}
	var n_time_for_output = []int{}

	// EXT_CODEの候補を必要数作成
	b_cnt := 0
	for b_cnt < B_CODE_COUNT {
		num := rand.Intn(300)
		if slices.Contains(n_time_list, num) == false {
			n_time_list = append(n_time_list, num)
			b_cnt++
		}
	}

	// 出力用のリスト作成
	list_cnt := 0
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for list_cnt < ROW_COUNT {
		random_index := r.Intn(B_CODE_COUNT)
		n_time_for_output = append(n_time_for_output, n_time_list[random_index])
		list_cnt++
	}

	ch <- n_time_for_output

	fmt.Println("create_n_time Done")
}

func create_telegram(wg *sync.WaitGroup, ch chan<- []string) {
	const time_format = "20060102150405000000"
	defer wg.Done()
	defer close(ch)
	var telegram_for_output = []string{}

	now := time.Now().Format(time_format) // YYYYMMDDHHMMSS + ミリ秒6桁
	telegram_for_output = append(telegram_for_output, now)

	telegram_cnt := 0
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for telegram_cnt < ROW_COUNT-1 {
		last_val := telegram_for_output[len(telegram_for_output)-1]
		if random_bool_val(r) {
			new_time, err := time.Parse(time_format, last_val)
			if err != nil {
				fmt.Println("Time Parse Error")
				break
			}
			t := new_time.Add(1000 * time.Millisecond)
			telegram_for_output = append(telegram_for_output, t.Format(time_format))
		} else {
			telegram_for_output = append(telegram_for_output, last_val)
		}
		telegram_cnt++
	}

	ch <- telegram_for_output

	fmt.Println("create_telegram Done")
}

func random_bool_val(r *rand.Rand) bool {
	val := r.Intn(2)
	if val == 0 {
		return true
	} else {
		return false
	}
}

func create_dataframe(ext_list []string, b_code []string, d_type []int, n_time []int, telegram []string) {
	// ROW_COUNTが大きいので、createの戻り値でサイズ制御できなさそう（channel element type too large (>64kB)）
	if len(ext_list) != ROW_COUNT {
		fmt.Println("ext_list size error : ", len(ext_list))
	}
	if len(b_code) != ROW_COUNT {
		fmt.Println("b_code size error : ", len(b_code))
	}
	if len(d_type) != ROW_COUNT {
		fmt.Println("d_type size error : ", len(d_type))
	}
	if len(n_time) != ROW_COUNT {
		fmt.Println("n_time size error : ", len(n_time))
	}
	if len(telegram) != ROW_COUNT {
		fmt.Println("telegram size error : ", len(telegram))
	}

	// 成型
	rows := make([]output_row, 0, ROW_COUNT)
	for i := 0; i < ROW_COUNT; i++ {
		rows = append(rows, output_row{
			EXT_CODE: ext_list[i],
			B_CODE:   b_code[i],
			D_TYPE:   d_type[i],
			N_TIME:   n_time[i],
			TELEGRAM: telegram[i],
		})
	}

	// fmt.Println(rows)

	f, err := os.Create(OUTPUT_FILENAME)
	if err != nil {
		fmt.Println("CSV create Error")
		return
	}
	defer f.Close()
	w := csv.NewWriter(f)
	defer w.Flush()

	// レコード

	if err := w.Write(HEADER); err != nil {
		return
	}
	for _, r := range rows {
		rec := []string{
			r.EXT_CODE,
			r.B_CODE,
			strconv.Itoa(r.D_TYPE),
			strconv.Itoa(r.N_TIME),
			r.TELEGRAM,
		}
		if err := w.Write(rec); err != nil {
			return
		}
	}

}

func launch[T any](wg *sync.WaitGroup, f func(*sync.WaitGroup, chan<- T)) <-chan T {
	ch := make(chan T, 1)
	wg.Add(1)
	go f(wg, ch)
	return ch
}

func create_csv() {
	currentTime := time.Now()
	var wg sync.WaitGroup

	chan_ext_code := launch(&wg, create_ext_code)
	chan_b_code := launch(&wg, create_b_code)
	chan_d_type := launch(&wg, create_d_type)
	chan_n_time := launch(&wg, create_n_time)
	chan_telegram := launch(&wg, create_telegram)

	list_ext_code := <-chan_ext_code
	list_b_code := <-chan_b_code
	list_d_type := <-chan_d_type
	list_n_time := <-chan_n_time
	list_telegram := <-chan_telegram

	wg.Wait()

	// csv作成
	file, err := os.Create(OUTPUT_FILENAME)
	if err != nil {
		fmt.Println("CSV create Error")
		return
	}
	defer file.Close()

	create_dataframe(list_ext_code, list_b_code, list_d_type, list_n_time, list_telegram)

	fmt.Println("Time : ", time.Since(currentTime))
}
