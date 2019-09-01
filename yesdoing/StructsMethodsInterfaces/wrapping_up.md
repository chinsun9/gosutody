# Interface, struct and method 정리

- 구조체 및 인터페이스
    - 구조체
        - 다른 언어들의 구조체, 클래스, 레코드 등과 비슷한 역할
        - 구조체를 이용하면 더 복잡한 자료형을 정의할 수 있다.
        - 필드들의 모음 or 묶음 === 구조체
        - 명명된 구성 요소들을 필드라고 불린다.
        - 배열이 서로 같은 자료형을 묶어 놓은 것이라면 구조체는 서로 다른 자료형의 자료들도 묶는다.
        - Table driven tests
            - anonymouse struct를 변수에 할당 할 수 있다.

                areaTests := []struct {
                        shape Shape
                        want  float64
                    }{
                        {Rectangle{12, 6}, 72.0},
                        {Circle{10}, 314.1592653589793},
                    }

        선언 

            type Task struct {
            	title string
            	done bool
            	due *time.Time
            }

    - 인터페이스
        - 메서드의 집합
        - 구현은 없고 메서드의 형태만 존재
        - 인터페이스의 메서드를 모두 구현하기만 하면 해당 인터페이스를 구현하는 것으로 취급(덕 타이핑) → 프로그램이 매우 유연해진다.
        - 빈 인터페이스 = 어떤 자료형도 받을 수 있는 와일드카드
        - Go 에서 Interface는 명시적이다.
        - interface 안에 선언된 메서드 명, 리턴 타입이 같아야지 같은 인터페이스라고 인식한다.
        - 인터페이스는 시스템의 다른 부분으로 부터 복잡성을 숨기기 위한 최고의 도구이다.
        - 우리의 테스트 케이스는 우리가 주입하고자 하는 것에 대해 정확히 무엇인지 알 필요가 없다. 오직 테스트 하고자 하는 것에 대해 어떻게 물을지만 생각하면 된다.
            ````go
            type InterfaceName interface {
            	FunctionName() returnType
            }
            ````
    - 메서드
        - 코드의 덩어리를 만든 다음에 그것을 호출하고 반환할 수 있는 구조를 서브루틴이라고 부른다.
        - 함수에 리시버(receiver)가 붙으면 메서드가 된다.
        - 자료형 T에 대하여 메서드를 호출할 때 이 자료형 T에 대한 리시버가 함수 이름, 즉 메서드 이름 앞에 붙는다.
        - `func (recv T) MethodName(p1, T1, p2 T2) R1`
        - 단순 자료형 메서드
            - 모든 명명된 자료형에서 메서드를 정의할 수 있다.
            - 기존에 존재하는 자료형을 쓸 수 있지만 여기에 이름을 붙여 추상화된 자료형을 사용하면 다른 것들과 구별되게 사용할 수 있다.
            ````go
            func (firstLetter ReceiverType) functionName () returnType {
            	//...logic
            }
            ````

### struct, method, interface가 Go언어 에서 왜 중요한가?
- 3가지 기능 모두 Go 환경에서의 서브루틴의 구현을 담당하기 때문에 중요하다.
- 소프트웨어를 디자인 할때 코드를 얼마나 알아보기 쉽게 추상화하는가가 중요한데 위 3가지 기능은 Go 환경에서의 코드의 추상화를 담당하며 이 기능을 통해 기존의 go 라이브러리를 응용한거나 사용할 수 있다.