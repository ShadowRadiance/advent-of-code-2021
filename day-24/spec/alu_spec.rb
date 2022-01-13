# frozen_string_literal: true

require './alu'
require 'stringio'

# rubocop:disable Metrics/BlockLength
RSpec.describe ALU do
  subject { ALU.new(input_stream: input_stream) }
  let(:input_stream) { StringIO.new('') }

  describe 'initial state' do
    it 'initializes all its registers to zero' do
      expect(subject.w).to eq(0)
      expect(subject.x).to eq(0)
      expect(subject.y).to eq(0)
      expect(subject.z).to eq(0)
    end
  end

  describe '#execute' do
    it 'splits the input string and calls perform' do
      expect(subject).to receive(:perform).with('aa', 'bb', 'cc')

      subject.execute('aa bb cc')
    end
  end

  describe '#perform' do
    context 'with bad parameters' do
      it 'complains about bad operations' do
        expect { subject.perform('not_an_operation', 'w', 0) }.to raise_error(
          ArgumentError,
          'not_an_operation is not a valid operation'
        )
      end

      it 'complains about invalid register names' do
        expect { subject.perform('add', 'q', 0) }.to raise_error(
          ArgumentError,
          'q is not a valid register'
        )
      end

      it 'complains about invalid second parameters' do
        expect { subject.perform('add', 'w', 'a') }.to raise_error(
          ArgumentError,
          'a is neither a valid register nor an integer'
        )

        expect { subject.perform('add', 'w', 19.5) }.to raise_error(
          ArgumentError,
          '19.5 is neither a valid register nor an integer'
        )
      end

      describe '#div' do
        it 'complains about div with a zero' do
          expect { subject.perform('div', 'w', 0) }.to raise_error(
            ArgumentError,
            '0 must not be zero'
          )
        end
      end

      describe '#mod' do
        context 'with a negative target register' do
          before { subject.perform('add', 'x', -5) }

          it 'complains about problems with mod' do
            expect { subject.perform('mod', 'x', 12) }.to raise_error(
              ArgumentError,
              'x (-5) must not be negative'
            )
          end
        end
        context 'with a positive target register' do
          before { subject.perform('add', 'x', 5) }

          it 'complains about problems with mod' do
            expect { subject.perform('mod', 'x', -5) }.to raise_error(
              ArgumentError,
              '-5 must be positive'
            )

            expect { subject.perform('mod', 'x', 0) }.to raise_error(
              ArgumentError,
              '0 must be positive'
            )
          end
        end
      end
    end

    context 'with missing data' do
      let(:input_stream) { StringIO.new('3') }
      it 'complains about missing data' do
        expect do
          subject.execute('inp w')
          subject.execute('inp x')
        end.to raise_error(EOFError)
      end
    end

    context 'with good parameters' do
      describe 'inp' do
        let(:input_stream) { StringIO.new('4321') }
        before do
          subject.execute('inp w')
          subject.execute('inp x')
          subject.execute('inp y')
          subject.execute('inp z')
        end
        it { is_expected.to have_attributes(w: 4, x: 3, y: 2, z: 1) }
      end

      describe 'add' do
        context 'with numeric second params' do
          let(:input_stream) { StringIO.new('4') }
          before do
            subject.execute('inp w')
            subject.execute('add w 2')
          end
          it { is_expected.to have_attributes(w: 6) }
        end
        context 'with register second params' do
          let(:input_stream) { StringIO.new('42') }
          before do
            subject.execute('inp w')
            subject.execute('inp x')
            subject.execute('add w x')
          end
          it { is_expected.to have_attributes(w: 6) }
        end
      end
      describe 'mul' do
        context 'with numeric second params' do
          let(:input_stream) { StringIO.new('4') }
          before do
            subject.execute('inp w')
            subject.execute('mul w 2')
          end
          it { is_expected.to have_attributes(w: 8) }
        end
        context 'with register second params' do
          let(:input_stream) { StringIO.new('42') }
          before do
            subject.execute('inp w')
            subject.execute('inp x')
            subject.execute('mul w x')
          end
          it { is_expected.to have_attributes(w: 8) }
        end
      end
      describe 'div' do
        context 'with numeric second params' do
          let(:input_stream) { StringIO.new('5') }
          before do
            subject.execute('inp w')
            subject.execute('div w 2')
          end
          it { is_expected.to have_attributes(w: 2) }
        end
        context 'with register second params' do
          let(:input_stream) { StringIO.new('52') }
          before do
            subject.execute('inp w')
            subject.execute('inp x')
            subject.execute('div w x')
          end
          it { is_expected.to have_attributes(w: 2) }
        end
      end
      describe 'mod' do
        context 'with numeric second params' do
          let(:input_stream) { StringIO.new('5') }
          before do
            subject.execute('inp w')
            subject.execute('mod w 3')
          end
          it { is_expected.to have_attributes(w: 2) }
        end
        context 'with register second params' do
          let(:input_stream) { StringIO.new('53') }
          before do
            subject.execute('inp w')
            subject.execute('inp x')
            subject.execute('mod w x')
          end
          it { is_expected.to have_attributes(w: 2) }
        end
      end
      describe 'eql' do
        context 'with numeric second params' do
          let(:input_stream) { StringIO.new('52') }
          before do
            subject.execute('inp w')
            subject.execute('eql w 5')
            subject.execute('inp z')
            subject.execute('eql z 5')
          end
          it { is_expected.to have_attributes(w: 1, z: 0) }
        end
        context 'with register second params' do
          let(:input_stream) { StringIO.new('55') }
          before do
            subject.execute('inp w')
            subject.execute('inp x')
            subject.execute('eql w x')
          end
          it { is_expected.to have_attributes(w: 1) }
        end
      end
    end
  end

  describe 'examples' do
    describe 'example 1' do
      let(:input_stream) { StringIO.new('7...') }
      before do
        # Here is an ALU program which takes an input number, negates it, and stores it in x:
        subject.execute('inp x')
        subject.execute('mul x -1')
      end
      it { is_expected.to have_attributes(x: -7) }
    end

    describe 'example 2' do
      let(:input_stream) { StringIO.new('39...') }
      before do
        # Here is an ALU program which takes two input numbers,
        # then sets z to 1 if the second input number is three times larger than the first input number,
        # or sets z to 0 otherwise:
        subject.execute('inp z')
        subject.execute('inp x')
        subject.execute('mul z 3')
        subject.execute('eql z x')
      end
      it { is_expected.to have_attributes(z: 1) }
    end

    describe 'example 3' do
      let(:input_stream) { StringIO.new('7...') }
      before do
        # Here is an ALU program which takes a non-negative integer as input,
        # converts it into binary, and stores the lowest (1's) bit in z,
        # the second-lowest (2's) bit in y, the third-lowest (4's) bit in x,
        # and the fourth-lowest (8's) bit in w:
        subject.execute('inp w')
        subject.execute('add z w')
        subject.execute('mod z 2')
        subject.execute('div w 2')
        subject.execute('add y w')
        subject.execute('mod y 2')
        subject.execute('div w 2')
        subject.execute('add x w')
        subject.execute('mod x 2')
        subject.execute('div w 2')
        subject.execute('mod w 2')
      end
      it { is_expected.to have_attributes(w: 0, x: 1, y: 1, z: 1) }
    end
  end
end
# rubocop:enable Metrics/BlockLength
