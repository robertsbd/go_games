func myPow(x float64, n int) float64 {

    output := 0;
    var _n float64 = 0;
    var _x float64 = 0;

    if n == 0 {return 1}

    if n < 0 { 
        _n = n * -1;
        _x = 1/x;
    } else {
        _n = n;
        _x = x;
    }

    for i := 0; i < _n - 1; i++ {
        output = output * _x;
    }
    
    return output;
}

    return outList.head; 
}

func PlainPow(x float64, n int) float64 {

    var output float64 = x;

    for i := 1; i < n; i++ {
        output = output * x;
    }

    return output;
}

func myPow(x float64, n int) float64 {

    var output float64 = 0;
    _n := 0;
    var _x float64 = 0;
    var sign float64 = 1;

    if n%2 == 1 && x < 0 {sign = -1}; 

    if n == 0 || x == 1 {return 1}

    if n == 1 {return x}

    if n < 0 { 
        _n = n * -1;
        _x = 1/x;
    } else {
        _n = n;
        _x = x;
    }

    var initial_input float64 = _x; 
    if _x < 0 {initial_input = _x*-1};

    if x == -1 {return initial_input * sign};

    // determine largest number that goes into n less than 500
    
    i := 0;

    for i = 500; _n%i != 0; i-- {}

    output = PlainPow(initial_input, i);
   
    output = PlainPow(output, int(_n / i));

    return output*sign;

}
