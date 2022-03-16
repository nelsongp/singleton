package singleton

import "sync"

/*instancia unica para que pueda ser accedida desde otro paquetes*/
var (
	p    *person
	once sync.Once
)

//el once se encarga de cuando el proceso es concurrente este solo sea ejecutado una sola vez por una instancia

func GetInstance() *person {
	//Al hacer el llamado en el do hacemos que la instanciacion del objeto se ejecutara una sola vez por cada vez que se le llame
	// eso si el do solo se ejecuta una sola vez por paquete y si existe otro do ya no la llamaragit
	once.Do(func() {
		p = &person{}
	})
	return p
}

/* dejamos la estrucutra privada para que no pueda ser accedida desde otros paquetes */
type person struct {
	name string
	age  int
	mux  sync.Mutex
}

/*exponemos los metodos que si podran ser invocados*/
func (p *person) SetName(n string) {
	p.mux.Lock()
	p.name = n
	p.mux.Unlock()
}

func (p *person) GetName() string {
	p.mux.Lock()
	defer p.mux.Unlock()
	return p.name
}

func (p *person) IncrementAge() {
	p.mux.Lock()
	p.age++
	p.mux.Unlock()
}

func (p *person) GetAge() int {
	p.mux.Lock()
	defer p.mux.Unlock()
	return p.age
}
