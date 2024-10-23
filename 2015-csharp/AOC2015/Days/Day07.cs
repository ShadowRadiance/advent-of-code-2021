using System.Text.RegularExpressions;

namespace AOC2015.Days;

public class Day07 : Day
{
    public Day07(string[] data) : base(data) { }

    private static readonly Regex _emitterRE = new(@"^(\d+)$");
    private static readonly Regex _andRE_Num = new(@"^(\d+) AND (\w+)$");
    private static readonly Regex _andRE = new(@"^(\w+) AND (\w+)$");
    private static readonly Regex _orRE = new(@"^(\w+) OR (\w+)$");
    private static readonly Regex _lshiftRE = new(@"^(\w+) LSHIFT (\d+)$");
    private static readonly Regex _rshiftRE = new(@"^(\w+) RSHIFT (\d+)$");
    private static readonly Regex _notRE = new(@"^NOT (\w+)$");
    private static readonly Regex _wireRE = new(@"^(\w+)$");


    class Wire(string name)
    {
        public string Name { get; } = name;
        public Gate? Input { get; set; } = null;
        private ushort _signal = 0;

        public ushort Signal
        {
            get
            {
                if (_signal == 0) _signal = EvaluateSignal();
                return _signal;
            }
        }

        public void Reset()
        {
            _signal = 0;
        }

        private ushort EvaluateSignal()
        {
            if (Input == null) return 0;

            return Input.Signal();
        }

        public override string ToString() => $"Wire({Name} <- {Input} ({_signal}))";
    }

    interface Gate { public ushort Signal(); };
    class Emitter(ushort value) : Gate
    {
        private readonly ushort _value = value;
        public ushort Signal() => _value;
        public override string ToString() => $"Emitter({_value})";
    }
    class AndGate(Wire left, Wire right) : Gate
    {
        private readonly Wire _left = left;
        private readonly Wire _right = right;
        public ushort Signal() => (ushort)(_left.Signal & _right.Signal);
        public override string ToString() => $"{_left.Name} AND {_right.Name})";
    }
    class OrGate(Wire left, Wire right) : Gate
    {
        private readonly Wire _left = left;
        private readonly Wire _right = right;
        public ushort Signal() => (ushort)(_left.Signal | _right.Signal);
        public override string ToString() => $"{_left.Name} AND {_right.Name})";
    }
    class LShiftGate(Wire input, ushort shift) : Gate
    {
        private readonly Wire _input = input;
        private readonly ushort _shift = shift;
        public ushort Signal() => (ushort)(_input.Signal << _shift);
        public override string ToString() => $"{_input.Name} LSHIFT {_shift})";
    }
    class RShiftGate(Wire input, ushort shift) : Gate
    {
        private readonly Wire _input = input;
        private readonly ushort _shift = shift;
        public ushort Signal() => (ushort)(_input.Signal >> _shift);
        public override string ToString() => $"{_input.Name} RSHIFT {_shift})";
    }
    class NotGate(Wire input) : Gate
    {
        private readonly Wire _input = input;
        public ushort Signal() => (ushort)~_input.Signal;
        public override string ToString() => $"NOT {_input.Name})";
    }
    class Passthrough(Wire input) : Gate
    {
        private readonly Wire _input = input;
        public ushort Signal() => _input.Signal;
        public override string ToString() => $"{_input.Name})";
    }

    class Solution
    {
        Dictionary<string, Wire> _wires = new();

        public Solution(string[] data)
        {
            foreach (string instruction in data)
            {
                string target = instruction.Split(" -> ").Last();
                if (!_wires.ContainsKey(target)) _wires.Add(target, new Wire(target));
            }

            foreach (string instruction in data)
            {
                string[] parts = instruction.Split(" -> ");
                string target = parts.Last();
                string input = parts.First();
                _wires[target].Input = BuildGate(input);
            }

            // foreach (var pair in wires)
            // {
            //     Console.WriteLine($"{pair.Key} <- {pair.Value}");
            // }
        }

        public ushort Signal(string wireName) => _wires[wireName].Signal;

        public void Reset()
        {
            foreach (Wire wire in _wires.Values)
            {
                wire.Reset();
            }
        }

        public void ReplaceInput(string wireName, string definition)
        {
            _wires[wireName].Input = BuildGate(definition);
        }

        private Gate BuildGate(string definition)
        {
            MatchCollection matches;

            matches = _emitterRE.Matches(definition);
            if (matches.Count == 1) return new Emitter(ushort.Parse(matches[0].Groups[1].Value));

            matches = _andRE_Num.Matches(definition);
            if (matches.Count == 1)
            {
                Wire dummy = new("dummy");
                dummy.Input = new Emitter(ushort.Parse(matches[0].Groups[1].Value));
                return new AndGate(dummy, _wires[matches[0].Groups[2].Value]);
            }

            matches = _andRE.Matches(definition);
            if (matches.Count == 1) { return new AndGate(_wires[matches[0].Groups[1].Value], _wires[matches[0].Groups[2].Value]); }

            matches = _orRE.Matches(definition);
            if (matches.Count == 1) return new OrGate(_wires[matches[0].Groups[1].Value], _wires[matches[0].Groups[2].Value]);

            matches = _lshiftRE.Matches(definition);
            if (matches.Count == 1) return new LShiftGate(_wires[matches[0].Groups[1].Value], ushort.Parse(matches[0].Groups[2].Value));

            matches = _rshiftRE.Matches(definition);
            if (matches.Count == 1) return new RShiftGate(_wires[matches[0].Groups[1].Value], ushort.Parse(matches[0].Groups[2].Value));

            matches = _notRE.Matches(definition);
            if (matches.Count == 1) return new NotGate(_wires[matches[0].Groups[1].Value]);

            matches = _wireRE.Matches(definition);
            if (matches.Count == 1) return new Passthrough(_wires[matches[0].Groups[1].Value]);

            throw new ArgumentException($"Cannot build gate from {definition}");
        }

    }

    public override string PartA()
    {
        /*
            Each wire has an identifier (some lowercase letters) and can carry a 16-bit signal (a number from 0 to 65535).
            A signal is provided to each wire by a gate, another wire, or some specific value.
            Each wire can only get a signal from one source, but can provide its signal to multiple destinations.
            A gate provides no signal until all of its inputs have a signal.
            The included instructions booklet describes how to connect the parts together:
                x AND y -> z means to connect wires x and y to an AND gate, and then connect its output to wire z.

            123 -> x means that the signal 123 is provided to wire x.
            x AND y -> z means that the bitwise AND of wire x and wire y is provided to wire z.
            p LSHIFT 2 -> q means that the value from wire p is left-shifted by 2 and then provided to wire q.
            NOT e -> f means that the bitwise complement of the value from wire e is provided to wire f.

            Other possible gates include OR (bitwise OR) and RSHIFT (right-shift).
            If, for some reason, you'd like to emulate the circuit instead, almost all programming languages
            (for example, C, JavaScript, or Python) provide operators for these gates.

            What signal is ultimately provided to wire "a"?
        */

        Solution solution = new(Data);
        return solution.Signal("a").ToString();
    }

    public override string PartB()
    {
        // Now, take the signal you got on wire a, override wire b to that signal,
        // and reset the other wires (including wire a).
        // What new signal is ultimately provided to wire "a"?

        Solution solution = new(Data);
        UInt16 firstRunSolution = solution.Signal("a");
        solution.Reset();
        solution.ReplaceInput("b", firstRunSolution.ToString());
        return solution.Signal("a").ToString();
    }
}
