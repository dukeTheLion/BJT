# Bipolar Junction Transistor Simple Calculator

## Objectives

- [x] Fixed Polarization Circuit (resistor)
- [x] Stable Emitter Polarization (resistor)
- [x] Voltage Divider Bias Circuit (resistor)
- [x] Fixed Polarization Circuit (current)
- [ ] Stable Emitter Polarization (current)
- [ ] Voltage Divider Bias Circuit (current)

## Command Line Format

| Fixed Polarization Circuit (resistor) | FR vcc rb rc beta |
|----|----|

| Stable Emitter Polarization (resistor) | SR vcc rb rc re beta |
|----|----|

| Voltage Divider Bias Circuit (resistor) | VDR vcc r1 r2 rc re beta |
|----|----|

| Fixed Polarization Circuit (current) | FC vcc ib ic vce |
|----|----|

## Output example

```text
- Fixed Polarization Circuit -

           vcc
            │
        ┌───┴───┐
        │       │  │IC
        │ │IB  rc  ▼    Vo
       rb └─►   ├────────►
        │      ┌┴┐ +
        └──────┤ │ VCE
               └┬┘ -
                │ E
               ─┴─
                -
                
  vcc =   12.0000 V
  rb =  240.0170 kΩ
  rc =    2.2000 kΩ
  β =   49.9150

  ib =   47.0800 uA
  ic =    2.3500 mA
  vbc =   -6.1300 V
  vce =    6.8300 V
  vb =  700.0000 mV
  vc =    6.8300 V
  ```